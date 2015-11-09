// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

// Example Usage
// 	url := "redis://localhost:6379/0"
// 	conn := bamboo.MakeConnFromUrl(url)
// 	queue := &RJQ{connection: conn, namespace: "MYAPP"}
type RJQ struct {
	namespace  string
	connection *Conn /* Redis connection */
}

/*
Conn An object representing the redis connection. This separate
object is used in order to abstract the redis client library.  User code only
needs to create a Conn object using the provided MakeConn and
MakeConnFromUrl functions and pass the result to RJQ.
*/
type Conn struct {
	client redis.Client
}

var MAX_RETRIES int = 3

type ConnectionError string

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("ConnectionError: %s", e)
}

/* MakeConn
 */
func MakeConn(host string, port int, pass string, db int) (conn *Conn) {
	client := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", host, port),
		Password:   pass,
		DB:         db,
		MaxRetries: MAX_RETRIES,
	})
	conn = &Conn{client: client}
	return conn
}

func MakeConnFromUrl(rawurl string) (conn *Conn) {
	urlp, err = url.Parse(rawurl)
	db, err := strconv.Atoi(strings.Split(urlp.Path, "/")[1])
	if err {
		panic
	}
	return makeConn(urlp.Host, urlp.Path)
}

func makeQueue(ns string, conn *Conn) (rjq *RJQ) {
	rjq = &RJQ{namespace: ns, connection: conn}
	return rjq
}

func (rjq *RJQ) Add(job Job) error {
	return nil
}

// TODO: Turn to variadic?
func (rjq *RJQ) AddMany(jobs []Job) error {
	for i := 0; i < len(jobs); i++ {
		err := rjq.Add(jobs[i])
		if err != nil {
			return err
		}
	}
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
