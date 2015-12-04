package bamboo

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func GenerateTestJobs(n int) (jobs []*Job) {
	for i := 0; i < n; i++ {
		jobid := fmt.Sprintf("job%d", i)
		data := make(map[string]string)
		data["a"] = "A"
		data["b"] = "B"
		payload, _ := json.Marshal(data)
		job := &Job{
			ID:          jobid,
			Priority:    5,
			Payload:     string(payload),
			ContentType: "application/json",
			Owner:       "",
			DateAdded:   time.Now().UTC().Unix(),
		}

		jobs = append(jobs, job)
	}
	return jobs
}

func TestJobSerialization(t *testing.T) {
	job := GenerateTestJobs(1)[0]
	jobarr := job.ToStringArray()
	// fmt.Println(jobarr)
	if !(len(jobarr) > 0) {
		t.Error("Job not serialized properly to array.")
	}
}
