// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

//go:generate go-bindata -pkg $GOPACKAGE -o scripts.go scripts/

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

const MAX_RETRIES int = 3
const SEP string = ":"

const ScriptNames = [...]string{
	"enqueue",
	"consume",
	"remove", // Can be accomplished with consume+ack
	"ack",
	"fail",
}

// Example Usage
// 	url := "redis://localhost:6379/0"
// 	conn := bamboo.MakeConnFromUrl(url)
// 	queue := &RJQ{Client: conn, Namespace: "MYAPP"}
type RJQ struct {
	Namespace string                  // Namespace within which all queues exist.
	Client    *Conn                   // Redis connection.
	Scripts   map[string]redis.Script // map of script name to redis.Script objects.
}

type ConnectionError string

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("ConnectionError: %s", e)
}

/*
Conn: An object representing the redis connection. This separate object is used
in order to abstract the redis client library.  User code only needs to create
a Conn object using the provided MakeConn and MakeConnFromUrl functions and
pass the result to RJQ.
*/
type Conn struct {
	client redis.Client
}

/* MakeConn
 */
func MakeConn(host string, port int, pass string, db int) (conn *Conn, error err) {
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", host, port),
		Password:   pass,
		DB:         db,
		MaxRetries: MAX_RETRIES,
	})
	conn = &Conn{client: client}
	return conn, nil
}

func MakeConnFromUrl(rawurl string) (conn *Conn, error err) {
	urlp, err = url.Parse(rawurl)
	db, err := strconv.Atoi(strings.Split(urlp.Path, "/")[1])
	if err {
		return nil, ConnectionError("Failed to parse url")
	}
	return MakeConn(urlp.Host, urlp.Path)
}

func MakeQueue(ns string, conn *Conn) (rjq *RJQ) {
	// Make the object
	rjq = &RJQ{Namespace: ns, Client: conn, Scripts: make(map[string]string)}

	// Load scripts
	for name := range ScriptNames {
		scriptSrc, err = Asset(fmt.Sprintf("scripts/%s.lua", name))
		if err != nil {
			panic(err) // Not loading the script means this program is incorrect.
		}
		script := rjq.Client.NewScript(scriptSrc)
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

func (rjq *RJQ) Add(job Job) error {
	queued_k := MakeKey(rjq.Namespace, "QUEUED")
	jobs_k := MakeKey(rjq.Namespace, "JOBS")
	keys := [...]string{queued_k, jobs_k}
	args := [...]string{}
	rjq.Client.EvalSha(sha, keys, args)

	// rjq.Connection.ZAdd(queued_k, score, job.id)
	// rjq.Connection.HMSet(jobs_k, (job.ToStringArray()))
	return nil
}

func (rjq *RJQ) GetOne() (Job, error) {
	var job Job
	return job, nil
}

func (rjq *RJQ) Get(string) (Job, error) {
	var job Job
	return job, nil
}

func (rjq *RJQ) Schedule(Job, time.Time) error {
	return nil
}

func (rjq *RJQ) Ack(Job) error {
	return nil
}

func (rjq *RJQ) Fail(Job) error {
	return nil
}

func (rjq *RJQ) Jobs(int) ([]Job, error) {
	var jobs []Job
	return jobs, nil
}

func (rjq *RJQ) Cancel(Job) error {
	return nil
}
