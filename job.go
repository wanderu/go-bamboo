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

func JobFromStringArray(arr []string) (*Job, error) {
	if len(arr)%2 == 1 {
		return nil, errors.New("Invalid number of key/value pairs in Job string array.")
	}

	job := &Job{}
	jtyp := reflect.TypeOf(*job) // deref ptr to get actual struct
	jval := reflect.ValueOf(job)

	// create a map from tag to field
	tagFieldMap := make(map[string]reflect.Value)
	for i := 0; i < jval.Elem().NumField(); i++ {
		tagFieldMap[jtyp.Field(i).Tag.Get("bamboo")] = jval.Elem().Field(i)
	}

	for i := 0; i < len(arr); i++ {
		key := strings.ToLower(arr[i]) // key is compared lowercase
		i++
		val := arr[i]

		// get field for tag
		field := tagFieldMap[key]

		switch field.Kind() {
		case reflect.Float64:
			valk, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Could not convert field %s to float.", key))
			}
			field.SetFloat(valk)
		case reflect.Int, reflect.Int32, reflect.Int64:
			valk, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Could not convert field %s to int.", key))
			}
			field.SetInt(valk)
		case reflect.String:
			field.SetString(val)
		}
	}

	return job, nil
}

func (job *Job) ToStringArray() []string {
	jtyp := reflect.TypeOf(*job) // deref ptr to get actual struct
	jval := reflect.ValueOf(*job)
	res := make([]string, int(jval.NumField())*2)
	for i := 0; i < jval.NumField(); i++ {
		field := jtyp.Field(i)
		tag := field.Tag.Get("bamboo")
		var val string
		switch field.Type.Kind() {
		case reflect.Float64:
			val = fmt.Sprintf("%f", jval.Field(i).Float())
		case reflect.Int, reflect.Int64:
			val = fmt.Sprintf("%d", jval.Field(i).Int())
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
