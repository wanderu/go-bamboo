// Use of this source code is governed by the MIT License,
// which can be found in the LICENSE file.

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
