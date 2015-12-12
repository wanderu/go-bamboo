/* Copyright 2015 Christopher Kirkos

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package bamboo

// Use `go generate` within the project directory after initializing
// the bamboo-scripts git submodule to generate the bamboo-scripts.go file.
//
//go:generate go-bindata -pkg $GOPACKAGE -o bamboo-scripts.go scripts/

import (
	"errors"
	"fmt"
	"gopkg.in/redis.v3"
	"math"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type QueueID string

const (
	QUEUED    QueueID = "QUEUED"
	SCHEDULED QueueID = "SCHEDULED"
	WORKING   QueueID = "WORKING"
	WORKERS   QueueID = "WORKERS"
	FAILED    QueueID = "FAILED"
)

const MAX_RETRIES int = 3
const SEP string = ":"

var ScriptNames = [...]string{
	"ack",
	"cancel",
	"consume",
	"enqueue",
	"fail",
	"maxfailed",
	"maxjobs",
	"recover",
	"test",
	// "remove", // Can be accomplished with consume+ack
}

type PSMsg struct {
	Msg string
	Err error
}

// Example Usage
// 	url := "redis://localhost:6379/0"
// 	conn := bamboo.MakeConnFromUrl(url)
// 	queue := &RJQ{Client: conn, Namespace: "MYAPP:QSET1"}
type RJQ struct {
	Namespace  string                   // Namespace within which all queues exist.
	Client     *redis.Client            // Redis connection.
	Scripts    map[string]*redis.Script // map of script name to redis.Script objects.
	WorkerName string
	JobExp     int // Seconds until a job expires
}

type ConnectionError string

func (e ConnectionError) Error() string {
	return fmt.Sprintf("ConnectionError: %s", e)
}

/* This time conversion loses precision in the nanoseconds due to the
conversion to/from float.
*/

func TimeToUnixTS(t time.Time) float64 {
	return float64(t.UnixNano()) / 1e+9
}

func UnixTSToTime(f float64) time.Time {
	return time.Unix(0, int64(f*1e+9))
}

/* MakeConn returns a redis connection object expected by RJQ.
 */
func MakeConn(host string, port int, pass string, db int64) (conn *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", host, port),
		Password:   pass,
		DB:         db,
		MaxRetries: MAX_RETRIES,
	})
	return client, nil
}

/* MakeConnFromUrl returns a redis connection object expected by RJQ.
 */
func MakeConnFromUrl(rawurl string) (conn *redis.Client, err error) {
	host, port := "127.0.0.1", 6379
	urlp, err := url.Parse(rawurl)
	host_port := strings.Split(urlp.Host, ":")
	if len(host_port) > 0 {
		host = host_port[0]
	}
	if len(host_port) > 1 {
		port, err = strconv.Atoi(host_port[1])
	}
	if err != nil {
		return nil, ConnectionError("Failed to parse port")
	}
	db, err := strconv.Atoi(strings.Split(urlp.Path, "/")[1])
	if err != nil {
		return nil, ConnectionError("Failed to parse db")
	}
	return MakeConn(host, port, "", int64(db))
}

func GenerateWorkerName() string {
	pid := os.Getpid()
	host, err := os.Hostname()
	// TODO: Warn hostname retrieval failed.
	if err != nil {
		host = fmt.Sprintf("%d", rand.Int31())
	}
	return fmt.Sprintf("%s-%d", host, pid) // TODO: Add random chars to end
}

func MakeQueue(ns string, conn *redis.Client) (rjq *RJQ) {
	// Make the object
	rjq = &RJQ{
		Namespace:  "{" + ns + "}",
		Client:     conn,
		Scripts:    make(map[string]*redis.Script),
		WorkerName: GenerateWorkerName(),
		JobExp:     90, // TODO: Parameterize
	}

	// Load scripts
	for _, name := range ScriptNames {
		scriptSrc, err := Asset(fmt.Sprintf("scripts/%s.lua", name))
		if err != nil {
			panic(err) // Not loading the script means this program is incorrect.
		}
		script := redis.NewScript(string(scriptSrc))
		res := script.Load(rjq.Client)
		if res.Err() != nil {
			panic(res.Err()) // If we can't load scripts, just fail.
		}
		rjq.Scripts[name] = script
	}

	// Set name with redis. Client.SetName(rjq.WorkerName)?

	return rjq
}

// Generate a key for the given namespace parts
// IE. MakeKey("a", "b", "c") -> "a:b:c"
func MakeKey(keys ...string) string {
	return strings.Join(keys, SEP)
}

// func MakeJobKey(ns string, jobid string) string {
// 	// {MY:NS:JOBS}:jobid
// 	return fmt.Sprintf("{%s}:%s", MakeKey(ns, "JOBS"), jobid)
// }

func (rjq RJQ) Subscribe() (chan PSMsg, error) {
	notify := make(chan PSMsg)
	pskey := MakeKey(rjq.Namespace, string(QUEUED))
	pubsub, err := rjq.Client.Subscribe(pskey)
	if err != nil {
		return notify, err
	}

	// Receive messages infinitely
	go func() {
		defer fmt.Println("anon func closed")
		defer pubsub.Close() // TODO: Determine if we need to close the PubSub conn
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				close(notify) // The sender closes
				return
			}
			notify <- PSMsg{msg.Payload, nil}
		}
	}()

	return notify, nil
}

func (rjq RJQ) Enqueue(job *Job, priority float64, queue QueueID, force string) error {
	if !(force == "1" || force == "0") {
		return errors.New("Invalid force paramter.")
	}
	// <ns>
	keys := []string{
		rjq.Namespace,
	}
	// <queue> <priority> <jobid> <force> <key> <val> [<key> <val> ...]
	args := []string{
		string(queue),
		fmt.Sprintf("%f", priority),
		job.ID,
		force,
	}
	args = append(args, job.ToStringArray()...)
	_, err := rjq.Scripts["enqueue"].EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rjq RJQ) Reschedule(job *Job, dt time.Time) error {
	return rjq.Enqueue(job, TimeToUnixTS(dt), SCHEDULED, "1")
}

func (rjq RJQ) Schedule(job *Job, dt time.Time) error {
	return rjq.Enqueue(job, TimeToUnixTS(dt), SCHEDULED, "0")
}

func (rjq RJQ) Requeue(job *Job, priority float64) error {
	return rjq.Enqueue(job, priority, QUEUED, "1")
}

func (rjq RJQ) Add(job *Job) error {
	// <ns> <kqueue> <kjob> <kqueued> <kscheduled> <kworking> <kfailed>
	return rjq.Enqueue(job, job.Priority, QUEUED, "0")
}

/* Consume returns the next available Job object from the queue.

Returns an error if arguments are invalid.
Returns an expected error if the max number of jobs has been reached for the
namespace.

Known error reply prefixes:
	MAXJOBS_REACHED: The maximum number of simultaneous jobs has been reached.
*/
func (rjq RJQ) Consume() (*Job, error) {
	// <ns>
	keys := []string{rjq.Namespace}
	// <client_name> <datetime> <job_id> <expires>
	args := []string{
		rjq.WorkerName,
		"",
		fmt.Sprintf("%d", time.Now().UTC().Unix()),
		fmt.Sprintf("%d", rjq.JobExp)}

	res := rjq.Scripts["consume"].EvalSha(rjq.Client, keys, args)
	if res.Err() != nil {
		return nil, res.Err()
	}

	// Convert the result into a string array and then into a Job object
	vals := res.Val().([]interface{})
	strs := make([]string, len(vals))
	for i, val := range vals {
		strs[i] = val.(string)
	}

	job, err := JobFromStringArray(strs)
	if err != nil {
		return nil, err
	}
	return job, nil
}

/* Get returns a Job object matching the specified jobid.
 */
func (rjq RJQ) Get(jobid string) (job *Job, err error) {
	job_key := MakeKey(rjq.Namespace, "JOBS", jobid)
	job_arr, err := rjq.Client.HGetAll(job_key).Result()
	if err != nil {
		return job, err
	}
	job, err = JobFromStringArray(job_arr)
	return job, nil
}

func (rjq RJQ) Ack(job *Job) error {
	keys := []string{rjq.Namespace}
	args := []string{job.ID}
	res := rjq.Scripts["ack"].EvalSha(rjq.Client, keys, args)
	if res.Err() != nil {
		// UNKNOWN_JOB_ID
		return res.Err()
	}
	return nil
}

func (rjq RJQ) Fail(job *Job, requeue_seconds int) error {
	keys := []string{rjq.Namespace}
	if requeue_seconds < 0 {
		// 0, 1, 4, 9, 16, 25 ... hours
		requeue_seconds = int(3600 * math.Pow(float64(job.Failures), 2))
	}
	args := []string{
		job.ID,
		fmt.Sprintf("%d", time.Now().UTC().Unix()),
		// seconds till requeue
		fmt.Sprintf("%d", requeue_seconds),
	}
	res := rjq.Scripts["fail"].EvalSha(rjq.Client, keys, args)
	if res.Err() != nil {
		// UNKNOWN_JOB_ID
		return res.Err()
	}
	return nil
}

func (rjq RJQ) Recover() ([]string, error) {
	keys := []string{rjq.Namespace}
	args := []string{
		fmt.Sprintf("%d", time.Now().UTC().Unix()),
		"3600",
	}
	res := rjq.Scripts["recover"].EvalSha(rjq.Client, keys, args)
	if res.Err() != nil {
		return nil, res.Err()
	}
	vals := res.Val().([]interface{})
	strs := make([]string, len(vals))
	for i, val := range vals {
		strs[i] = val.(string)
	}
	return strs, res.Err()
}

type JobResult struct {
	Job   *Job
	Error error
}

// Peek returns a channel Job objects.
//   queue: Queue from which to list Jobs.
func (rjq RJQ) Peek(queue QueueID) chan JobResult {
	srch := ZScanIter(rjq.Client, MakeKey(rjq.Namespace, string(queue)), "")
	jobch := make(chan JobResult)

	go func() {
		defer close(jobch)
		for sr := range srch {
			// Check for an error coming from the StringResult channel
			// Turn it into a JobResult
			if sr.Error != nil {
				jobch <- JobResult{nil, sr.Error}
				return
			}

			// TODO: Check sr.Result for valid job data
			//       Make Get return an error if the job does not exist.
			job, err := rjq.Get(sr.Result)

			// Check for an error with fetching the Job itself
			if err != nil {
				jobch <- JobResult{nil, err}
				return
			}

			jobch <- JobResult{job, nil}
		}
	}()

	return jobch
}

func (rjq RJQ) CancelById(jobid string) error {
	keys := []string{rjq.Namespace}
	args := []string{jobid}
	_, err := rjq.Scripts["cancel"].EvalSha(rjq.Client, keys, args).Result()
	return err
}

func (rjq RJQ) Cancel(job *Job) error {
	return rjq.CancelById(job.ID)
}

func (rjq RJQ) SetMaxFailed(n int) (int, error) {
	keys := []string{rjq.Namespace}
	args := []string{fmt.Sprintf("%d", n)}
	val, err := rjq.Scripts["maxfailed"].EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return 0, err
	}
	maxfailed, err := strconv.Atoi(val.(string))
	return maxfailed, err
}

func (rjq RJQ) GetMaxFailed(n int) (int, error) {
	keys := []string{rjq.Namespace}
	args := []string{""}
	val, err := rjq.Scripts["maxfailed"].EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return 0, err
	}
	maxfailed, err := strconv.Atoi(val.(string))
	return maxfailed, err
}

func (rjq RJQ) SetMaxJobs(n int) (int, error) {
	keys := []string{rjq.Namespace}
	args := []string{fmt.Sprintf("%d", n)}
	val, err := rjq.Scripts["maxjobs"].EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return 0, err
	}
	maxfailed, err := strconv.Atoi(val.(string))
	return maxfailed, err
}

func (rjq RJQ) GetMaxJobs(n int) (int, error) {
	keys := []string{rjq.Namespace}
	args := []string{""}
	val, err := rjq.Scripts["maxjobs"].EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return 0, err
	}
	maxfailed, err := strconv.Atoi(val.(string))
	return maxfailed, err
}

func (rjq RJQ) Test() error {
	fmt.Println("RJQ.Test calling lua test script")
	val, err := rjq.Scripts["test"].EvalSha(rjq.Client, []string{}, []string{}).Result()
	if err != nil {
		fmt.Println("RJQ.Test lua script error", err)
	} else {
		fmt.Println("RJQ.Test lua script result", val)
		// fmt.Printf("%#T\n", val)
	}
	return err
}
