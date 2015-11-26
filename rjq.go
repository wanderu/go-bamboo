// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

// Use `go generate` within the project directory after setting the
// BAMBOO_SCRIPTS_PFX and BAMBOO_SCRIPTS environment variables to generate the
// bamboo-scripts.go file.
//
//go:generate go-bindata -prefix $BAMBOO_SCRIPTS_PFX -pkg $GOPACKAGE -o bamboo-scripts.go $BAMBOO_SCRIPTS

import (
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
	return fmt.Sprintf("%s-%d", host, pid)
}

func MakeQueue(ns string, conn *redis.Client) (rjq *RJQ) {
	// Make the object
	rjq = &RJQ{
		Namespace:  ns,
		Client:     conn,
		Scripts:    make(map[string]*redis.Script),
		WorkerName: GenerateWorkerName(),
		JobExp:     90, // TODO: Parameterize
	}

	// Load scripts
	for _, name := range ScriptNames {
		scriptSrc, err := Asset(fmt.Sprintf("bamboo-scripts/%s.lua", name))
		if err != nil {
			panic(err) // Not loading the script means this program is incorrect.
		}
		script := redis.NewScript(string(scriptSrc))
		script.Load(rjq.Client)
		rjq.Scripts[name] = script
	}

	return rjq
}

// Generate a key for the given namespace parts
// IE. MakeKey("a", "b", "c") -> "a:b:c"
func MakeKey(keys ...string) string {
	return strings.Join(keys, SEP)
}

func (rjq RJQ) Add(job *Job) error {
	keys := []string{rjq.Namespace}
	args := []string{fmt.Sprintf("%f", job.Priority), job.JobID}
	args = append(args, job.ToStringArray()...)
	val, err := rjq.Scripts["enqueue"].EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return err
	}
	// val should be 0 if the job already existed and 1 if the job was enqueued
	// as expected.
	if val == 0 {
		// TODO: log that item already existed on queue
	} else {
		// TODO: log successful addition of job
	}
	return nil
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

func (rjq RJQ) Schedule(*Job, time.Time) error {
	return nil
}

func (rjq RJQ) Ack(job *Job) error {
	keys := []string{rjq.Namespace}
	args := []string{job.JobID}
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
		requeue_seconds = int(3600 * math.Pow(float64(job.Failures), 2))
	}
	args := []string{
		job.JobID,
		fmt.Sprintf("%d", time.Now().UTC().Unix()),
		// seconds till requeue
		fmt.Sprintf("%f", requeue_seconds),
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

// Peek returns the top n jobs from the queue.
func (rjq RJQ) Peek(n int, queue QueueID) (jobs []*Job, err error) {
	ids, err := rjq.Client.ZRange(string(queue), 0, int64(n)).Result()
	for _, id := range ids {
		job, err := rjq.Get(id)
		if err != nil {
			continue
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (rjq RJQ) Cancel(job *Job) error {
	keys := []string{rjq.Namespace}
	args := []string{job.JobID}
	res := rjq.Scripts["cancel"].EvalSha(rjq.Client, keys, args)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
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
	val, err := rjq.Scripts["test"].EvalSha(rjq.Client, []string{}, []string{}).Result()
	// if strings.HasPrefix(err.Error(), "MAXJOBS") {
	// 	fmt.Println("Found MAXJOBS Error")
	// }
	fmt.Println("Error", err)
	fmt.Println("Result", val)
	fmt.Printf("%#T\n", val)
	return err
}
