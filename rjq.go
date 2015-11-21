// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

//go:generate go-bindata -pkg $GOPACKAGE -o scripts.go scripts/

import (
	"fmt"
	"gopkg.in/redis.v3"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const MAX_RETRIES int = 3
const SEP string = ":"

var ScriptNames = [...]string{
	"enqueue",
	"consume",
	// "ack",
	// "fail",
	// "test",
	// "remove", // Can be accomplished with consume+ack
}

// Example Usage
// 	url := "redis://localhost:6379/0"
// 	conn := bamboo.MakeConnFromUrl(url)
// 	queue := &RJQ{Client: conn, Namespace: "MYAPP:QSET1"}
// Namespaces are colon-separated string segments.
type RJQ struct {
	Namespace string                   // Namespace within which all queues exist.
	Client    *redis.Client            // Redis connection.
	Scripts   map[string]*redis.Script // map of script name to redis.Script objects.
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

func MakeQueue(ns string, conn *redis.Client) (rjq *RJQ) {
	// Make the object
	rjq = &RJQ{Namespace: ns, Client: conn, Scripts: make(map[string]*redis.Script)}

	// Load scripts
	for _, name := range ScriptNames {
		scriptSrc, err := Asset(fmt.Sprintf("scripts/%s.lua", name))
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

func (rjq *RJQ) Add(job *Job) error {
	keys := []string{rjq.Namespace}
	args := []string{fmt.Sprintf("%d", job.Priority), job.JobID}
	args = append(args, job.ToStringArray()...)
	script := rjq.Scripts["enqueue"]
	val, err := script.EvalSha(rjq.Client, keys, args).Result()
	if err != nil {
		return err
	}
	if val == 0 {
		// TODO: log that item already existed on queue
	} else {
		// TODO: log successful addition of job
	}
	return nil
}

/* GetOne returns the next available Job object from the queue.

Returns an error if arguments are invalid.
Returns an expected error if the max number of jobs has been reached for the
namespace.

Known error reply prefixes:
	MAXJOBS: The maximum number of simultaneous jobs has been reached.
*/
func (rjq *RJQ) GetOne() (*Job, error) {
	keys := []string{rjq.Namespace}
	args := []string{"", "", "", ""}
	script := rjq.Scripts["consume"]

	res := script.EvalSha(rjq.Client, keys, args)
	if res.Err() != nil {
		fmt.Println(res.Err())
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
func (rjq *RJQ) Get(jobid string) (*Job, error) {
	var job *Job
	return job, nil
}

func (rjq *RJQ) Schedule(*Job, time.Time) error {
	return nil
}

func (rjq *RJQ) Ack(*Job) error {
	return nil
}

func (rjq *RJQ) Fail(*Job) error {
	return nil
}

func (rjq *RJQ) Jobs(int) ([]*Job, error) {
	var jobs []*Job
	return jobs, nil
}

func (rjq *RJQ) Cancel(*Job) error {
	return nil
}

/*
func (rjq *RJQ) Test() error {
	script := rjq.Scripts["test"]
	val, err := script.EvalSha(rjq.Client, []string{}, []string{}).Result()
	if strings.HasPrefix(err.Error(), "MAXJOBS") {
		fmt.Println("Found MAXJOBS Error")
	}
	fmt.Println("Error", err)
	fmt.Println("Result", val)
	return err
}
*/

// queue_k := MakeKey(rjq.Namespace, "QUEUED")
// jobs_k := MakeKey(rjq.Namespace, "JOBS", job.JobID)
// queue_k := MakeKey(rjq.Namespace, "QUEUED")
// working_k := MakeKey(rjq.Namespace, "QUEUED")
// job_ns := rjq.Namespace
// maxjobs_k := MakeKey(rjq.Namespace, "MAXJOBS")
// keys := []string{queue_k, working_k, job_ns, maxjobs_k}
