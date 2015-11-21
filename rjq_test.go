// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMakeKey(t *testing.T) {
	key := MakeKey("a", "b", "c")
	expected := "a:b:c"
	if key != expected {
		t.Errorf("MakeKey() Returned %s Expected %s", key, expected)
	}
}

func TestTest(t *testing.T) {
	ns := "TEST"
	conn, _ := MakeConn("localhost", 6379, "", 0)
	rjq := MakeQueue(ns, conn)
	rjq.Test()
}

func TestAdd(t *testing.T) {
	// 1. Establish Connection
	ns := "TEST"
	conn, err := MakeConn("localhost", 6379, "", 0)
	rjq := MakeQueue(ns, conn)

	// 2. Make and add Job
	data := make(map[string]string)
	data["a"] = "A"
	data["b"] = "B"
	payload, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	job1id := "job1"
	job := &Job{
		JobID:       job1id,
		Priority:    5,
		Payload:     string(payload),
		ContentType: "application/json",
		Worker:      GenerateWorkerName(),
		DateAdded:   time.Now().UTC().Unix(),
	}

	job_key := MakeKey(rjq.Namespace, "JOBS", job1id)

	// Make sure it doesn't exist yet before adding
	_ = rjq.Client.Del(job_key)

	rjq.Add(job)

	// 3. Test the stored job data directly in Redis
	job_arr, err := rjq.Client.HGetAll(job_key).Result()
	if err != nil {
		t.Error(err)
	}

	// 3b. Test conversion to Job object
	job2, err := JobFromStringArray(job_arr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(job2)

	if !reflect.DeepEqual(job, job2) {
		t.Error("Job data doesn't match.")
	}

	// 4. Consume Job
	job3, err := rjq.GetOne()
	fmt.Println(job3)

	// 5. Verify Job Contents
	if !reflect.DeepEqual(job, job3) {
		t.Error("Job data doesn't match.")
	}

	// 6. Test that the Job is on the right queue
	// 7. Ack the Job
	// 8. Test that the Job does not exist on any queue
}
