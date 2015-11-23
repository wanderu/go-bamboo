// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

package bamboo

import (
	"encoding/json"
	"fmt"
	// "reflect"
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
	// fmt.Println(rjq)
	rjq.Test()
}

func TestMaxFailed(t *testing.T) {
	ns := "TEST"
	conn, _ := MakeConn("localhost", 6379, "", 0)
	rjq := MakeQueue(ns, conn)
	kmaxfailed := MakeKey(ns, "MAXFAILED")
	_ = rjq.Client.Del(kmaxfailed)
	maxfailed := 5
	n, err := rjq.SetMaxFailed(maxfailed)
	if err != nil {
		t.Error(err)
	}
	if n != maxfailed {
		t.Error(fmt.Sprintf("%d != %d", n, maxfailed))
	}
	_ = rjq.Client.Del(kmaxfailed)
}

func TestMaxJobs(t *testing.T) {
	ns := "TEST"
	conn, _ := MakeConn("localhost", 6379, "", 0)
	rjq := MakeQueue(ns, conn)
	kmaxjobs := MakeKey(ns, "MAXJOBS")
	_ = rjq.Client.Del(kmaxjobs)
	maxjobs := 5
	n, err := rjq.SetMaxFailed(maxjobs)
	if err != nil {
		t.Error(err)
	}
	if n != maxjobs {
		t.Error(fmt.Sprintf("%d != %d", n, maxjobs))
	}
	_ = rjq.Client.Del(kmaxjobs)
}

// ComprareJobs compares a subset of job fields that should not change
// over the lifetime of the job.
func CompareJobs(a *Job, b *Job) bool {
	return a.Priority == b.Priority &&
		a.JobID == b.JobID &&
		a.Payload == b.Payload &&
		a.DateAdded == b.DateAdded &&
		a.ContentType == b.ContentType &&
		a.Encoding == b.Encoding
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
		Owner:       "",
		DateAdded:   time.Now().UTC().Unix(),
	}

	job_key := MakeKey(rjq.Namespace, "JOBS", job1id)

	// Make sure it doesn't exist yet before adding
	_ = rjq.Client.Del(job_key)

	kqueued := MakeKey(rjq.Namespace, "QUEUED")
	kworking := MakeKey(rjq.Namespace, "WORKING")
	kscheduled := MakeKey(rjq.Namespace, "SCHEDULED")
	kfailed := MakeKey(rjq.Namespace, "FAILED")
	// fmt.Println("kqueued: ", kqueued)
	// fmt.Println("kworking: ", kworking)
	_ = rjq.Client.ZRem(kqueued, job1id)
	_ = rjq.Client.ZRem(kworking, job1id)

	err = rjq.Add(job)
	if err != nil {
		t.Error(err)
	}

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
	if job2.State != "enqueued" {
		t.Error("State is not enqueued. State: ", job2.State)
	}
	// fmt.Println(job)
	// fmt.Println(job2)

	if !CompareJobs(job, job2) {
		t.Error("Job data doesn't match.")
	}

	// 4. Consume Job
	job3, err := rjq.GetOne()
	// fmt.Println(job3)
	if err != nil {
		fmt.Println(err)
	}

	// 5. Verify Job Contents
	// if !reflect.DeepEqual(job, job3) {
	if !CompareJobs(job, job3) {
		t.Error("Job data doesn't match.")
	}
	// Check owner ID
	if job3.Owner != rjq.WorkerName {
		t.Error("Job Owner doesn't match WorkerName.")
	}

	// 6. Test that the Job is on the right queue
	res, err := rjq.Client.ZScore(kworking, job1id).Result()
	// fmt.Println("Job score", res)

	if res != job.Priority {
		t.Error("Job Priority does not match.", res)
	}

	// 7. Ack the Job
	err = rjq.Ack(job3)
	if err != nil {
		t.Error(err)
	}

	// 8. Test that the Job does not exist on any queue
	res, err = rjq.Client.ZScore(kqueued, job1id).Result()
	if res > 0 {
		t.Error("Job still in queue.", job1id, "Score", res)
	}
	res, err = rjq.Client.ZScore(kworking, job1id).Result()
	if res > 0 {
		t.Error("Job still in working queue.", job1id, "Score", res)
	}

	// Test fail case
	err = rjq.Add(job)
	if err != nil {
		t.Error(err)
	}
	err = rjq.Fail(job)
	job4, err := rjq.GetOne()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job4)
	res, err = rjq.Client.ZScore(kscheduled, job1id).Result()
	// should not be on the scheduled queue
	res, err = rjq.Client.ZScore(kfailed, job1id).Result()
	// should be on the failed queue

}
