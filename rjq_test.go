// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

import (
	"encoding/json"
	"fmt"
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
		JobID: job1id, Priority: 5, Payload: string(payload),
		ContentType: "application/json",
		Worker:      GenerateWorkerName(),
		DateAdded:   time.Now().UTC().Unix(),
	}
	rjq.Add(job)

	job_hm := rjq.Client.HMGet(
		MakeKey(rjq.Namespace, "JOBS", job1id), "*")
	if job_hm.Err() != nil {
		t.Error(job_hm.Err())
	}
	job_map := job_hm.Val()
	fmt.Println(job_map)
	// if job_map["a"] != "A" {
	// 	t.Error("Bad data")
	// }

	// 3. Verify Count
	// 4. Consume Job
	// 5. Verify Job Contents
}
