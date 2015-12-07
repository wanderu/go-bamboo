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
	"fmt"
	"time"
)

type Queue interface {
	Add(Job) error
	GetOne() (Job, error)
	Get(string) (Job, error)
	Schedule(Job, time.Time) error
	Ack(Job) error
	Fail(Job) error
	Jobs(int) ([]Job, error)
	Cancel(Job) error
}

type QueueError string

func (e *QueueError) Error() string {
	return fmt.Sprintf("QueueError: %s", e)
}
