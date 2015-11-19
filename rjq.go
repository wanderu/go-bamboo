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
	// "consume",
	// "remove", // Can be accomplished with consume+ack
	// "ack",
	// "fail",
}

// Example Usage
// 	url := "redis://localhost:6379/0"
// 	conn := bamboo.MakeConnFromUrl(url)
// 	queue := &RJQ{Client: conn, Namespace: "MYAPP"}
type RJQ struct {
	Namespace string // Namespace within which all queues exist.
	// Client    *Conn               // Redis connection.
	Client  *redis.Client            // Redis connection.
	Scripts map[string]*redis.Script // map of script name to redis.Script objects.
}

type ConnectionError string

func (e ConnectionError) Error() string {
	return fmt.Sprintf("ConnectionError: %s", e)
}

/*
Conn: An object representing the redis connection. This separate object is used
in order to abstract the redis client library.  User code only needs to create
a Conn object using the provided MakeConn and MakeConnFromUrl functions and
pass the result to RJQ.
*/
// type Conn struct {
// 	client redis.Client
// }

/* MakeConn
 */
func MakeConn(host string, port int, pass string, db int64) (conn *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", host, port),
		Password:   pass,
		DB:         db,
		MaxRetries: MAX_RETRIES,
	})
	// conn = &Conn{client: client}
	return client, nil
}

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
	queue_k := MakeKey(rjq.Namespace, "QUEUED")
	jobs_k := MakeKey(rjq.Namespace, "JOBS", job.JobID)
	keys := []string{queue_k, jobs_k}
	args := []string{fmt.Sprintf("%d", job.Priority), job.JobID}
	for _, v := range job.ToStringArray() {
		args = append(args, v)
	}
	script := rjq.Scripts["enqueue"]
	// fmt.Println("calling enqueue")
	// fmt.Println(args)
	res := script.Eval(rjq.Client, keys, args)
	fmt.Println("---")
	fmt.Println(res)
	fmt.Println("---")
	if res.Err() != nil {
		return res.Err()
	}
	// rjq.Connection.ZAdd(queued_k, score, job.id)
	// rjq.Connection.HMSet(jobs_k, (job.ToStringArray()))
	return nil
}

func (rjq *RJQ) GetOne() (*Job, error) {
	var job *Job
	return job, nil
}

func (rjq *RJQ) Get(string) (*Job, error) {
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
