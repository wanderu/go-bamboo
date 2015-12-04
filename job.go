// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

// time.RFC3339 // "2006-01-02T15:04:05Z07:00"
const (
	MDY = "01/02/2006"
)

func itos(i int) string {
	return fmt.Sprintf("%d", i)
}

func f64tos(f float64) string {
	return fmt.Sprintf("%f", f)
}

func i64tos(i int64) string {
	return fmt.Sprintf("%d", i)
}

// Job is an object representing the item stored in the queue.
// Volatile paramters change over the life time of the Job.
type Job struct {
	// Required - Set by user
	ID       string  `bamboo:"id"`       // Unique per queue instance.
	Priority float64 `bamboo:"priority"` // Used by priority queues for ordering.
	Payload  string  `bamboo:"payload"`  // Byte string.
	// Set by queuing system
	DateAdded  int64  `bamboo:"created"`  // Unix Timestamp of the job creation date.
	Failures   int    `bamboo:"failures"` // Number of failures so far. (Volatile)
	DateFailed int64  `bamboo:"failed"`   // Unix Timestamp of the last failure. (Volatile)
	Consumed   int64  `bamboo:"consumed"` // Unix Timestamp of the date this job was consumed. (Volatile)
	Owner      string `bamboo:"owner"`    // Name of the worker process. (Volatile)
	// Not required - Set by user
	ContentType string `bamboo:"contenttype"` // ContentType of the payload IE. 'application/json'.
	Encoding    string `bamboo:"encoding"`    // Encoding of the payload.
	State       string `bamboo:"state"`       // Job state. "enqueued", "working", "failed". (Volatile)
}

// func JobFromStringArrayTags(arr []string) (*Job, error) {
// 	job := &Job{}
// 	var err interface{}

// 	field, ok := Job.FieldByName()
// }

func JobFromStringArray(arr []string) (*Job, error) {
	job := &Job{}
	var err interface{}
	for i := 0; i < len(arr); i++ {
		key := strings.ToLower(arr[i]) // key is compared lowercase
		i++
		val := arr[i]
		switch key {
		case "priority":
			job.Priority, err = strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, errors.New("Could not convert priority to int.")
			}
		case "id":
			job.ID = val
		case "payload":
			job.Payload = val
		case "failures":
			job.Failures, err = strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("Cound not convert failures to int.")
			}
		case "dateadded":
			job.DateAdded, err = strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, errors.New("Cound not convert dateadded to int.")
			}
		case "datefailed":
			job.DateFailed, err = strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, errors.New("Cound not convert datefailed to int.")
			}
		case "consumed":
			job.Consumed, err = strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, errors.New("Cound not convert consumed to int.")
			}
		case "owner":
			job.Owner = val
		case "contenttype":
			job.ContentType = val
		case "encoding":
			job.Encoding = val
		case "state":
			job.State = val
		default:
			return nil, errors.New(fmt.Sprintf("Invalid key: %s", key))
		}
	}
	return job, nil
}

func (job *Job) ToStringArray() []string {
	jtyp := reflect.TypeOf((*job)) // deref ptr
	jval := reflect.ValueOf((*job))
	res := make([]string, int(jval.NumField())*2)
	for i := 0; i < jval.NumField(); i++ {
		field := jtyp.Field(i)
		tag := field.Tag.Get("bamboo")
		var val string
		switch field.Type.Kind() {
		case reflect.Float64:
			val = f64tos(jval.Field(i).Float())
		case reflect.Int:
			val = i64tos(jval.Field(i).Int())
		case reflect.Int64:
			val = i64tos(jval.Field(i).Int())
		case reflect.String:
			val = jval.Field(i).String()
		default:
			panic("Invalid type")
		}
		res[i*2] = tag
		res[(i*2)+1] = val
	}
	return res
}

func (job *Job) ToStringArrayManual() []string {
	return []string{
		"priority", f64tos(job.Priority),
		"id", job.ID,
		"payload", job.Payload,
		"failures", itos(job.Failures),
		"dateadded", i64tos(job.DateAdded),
		"datefailed", i64tos(job.DateFailed),
		"consumed", i64tos(job.Consumed),
		"owner", job.Owner,
		"contenttype", job.ContentType,
		"encoding", job.Encoding,
		"state", job.State,
	}
}

// MakeJobID generates a random Job ID.
func MakeJobID() string {
	return fmt.Sprintf("%d", rand.Int63())
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
