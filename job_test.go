/* Copyright 2015 Christopher Kirkos

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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

// ComprareJobs compares a subset of job fields that should not change
// over the lifetime of the job.
func CompareJobs(a *Job, b *Job) bool {
	return a.Priority == b.Priority &&
		a.ID == b.ID &&
		a.Payload == b.Payload &&
		a.DateAdded == b.DateAdded &&
		a.ContentType == b.ContentType &&
		a.Encoding == b.Encoding
}

func TestJobSerialization(t *testing.T) {
	job := GenerateTestJobs(1)[0]
	jobarr := job.ToStringArray()

	if !(len(jobarr) > 0) {
		t.Fatal("Job not serialized properly to array.")
	}

	job2, err := JobFromStringArray(jobarr)

	if err != nil {
		t.Fatal("Error in JobFromStringArray", err)
	}

	// Ensure it's not empty
	if CompareJobs(job, &Job{}) {
		t.Fatal("Empty job. Invalid.")
	}

	// Ensure it's the same as the input Job
	if !CompareJobs(job, job2) {
		t.Fatal("Jobs don't match")
	}
}
