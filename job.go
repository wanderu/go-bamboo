// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

import (
	"fmt"
	"math/rand"
	"os"
	// "strconv"
	// "time"
)

// time.RFC3339 // "2006-01-02T15:04:05Z07:00"
const (
	MDY = "01/02/2006"
)

func itos(i int) string {
	return fmt.Sprintf("%d", i)
}

func i64tos(i int64) string {
	return fmt.Sprintf("%d", i)
}

// Job is an object representing the item stored in the queue.
type Job struct {
	// Required - Set by user
	Priority int    // Used by priority queues for ordering.
	JobID    string // Unique per queue instance.
	Payload  string // Byte string.
	// Set by queuing system
	Failures   int    // Number of failures so far.
	DateAdded  int64  // Unix Timestamp of the job creation date.
	DateFailed int64  // Unix Timestamp of the last failure.
	Worker     string // Name of the worker process.
	// Not required - Set by user
	ContentType string // ContentType of the payload IE. 'application/json'.
	Encoding    string // Encoding of the payload.
}

func (job *Job) ToStringArray() []string {
	return []string{
		"priority", itos(job.Priority),
		"jobid", job.JobID,
		"payload", job.Payload,
		"failures", itos(job.Failures),
		"dateadded", i64tos(job.DateAdded),
		"datefailed", i64tos(job.DateFailed),
		"worker", job.Worker,
		"contenttype", job.ContentType,
		"encoding", job.Encoding,
	}
}

// MakeJobID generates a random Job ID.
func MakeJobID() string {
	return fmt.Sprintf("%d", rand.Int63())
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

/*
Example Use

import "encoding/json"

func MakeJob(id string, priority int) (job *Job) {
	job = &Job{JobID: id, Priority: priority, Failures: 0, ContentType: "", Encoding: "", Worker: ""}
	return job
}

func (job *Job) SetPayload(payload string, contentType string, encoding string) {
	job.Payload = payload
	job.ContentType = contentType
	job.Encoding = encoding
}

func (job *Job) SetTextPayload(payload string) {
	job.setPayload(payload, "text/plain", "")
}

func (job *Job) SetJsonPayload(payload interface{}) error {
	jsonStr, err := json.Marshall(payload)
	if err != nil {
		return err
	}
	job.setPayload(jsonStr, "application/json", "")
}

func (job *Job) DecodeJsonPayload() error {
	return json.Unmarshal(job.Payload, &Job.Payload)
}

// If we have a byte array, turn it into a string.
func (job *Job) SetBytePayload(payload []byte) {
	job.SetPayload(string(payload[:]), "application/octet-stream", "")
}

import "time"

// NowTimestamp returns the current UTC Unix Timestamp in seconds.
func NowTimestamp() int64 {
	return time.Now().UTC().Unix()
}

// FromUnixTimestamp returns a time.Time object given a valid Unix timestamp in
// seconds.
func FromUnixTimestamp(seconds int64) time.Time {
	return time.Unix(seconds, 0)
}
*/
