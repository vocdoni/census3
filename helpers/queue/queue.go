// queue package abstracts a structure for enqueuing processes that running in
// the background, updating the status of those processes and dequeuing them.
// It ensures that processes are enqueued, dequeued and update thread safely.
package queue

import (
	"sync"

	"go.vocdoni.io/dvote/util"
)

// QueueIDLen constant defines the fixed length of any queue item
const QueueIDLen = 20

// QueueItem struct defines a single queue item, including if it is done or not,
// it it throw any error during its execution, and a flexsible data map variable
// to store resulting information of the enqueued process execution.
type QueueItem struct {
	Done     bool    `json:"done"`
	Error    error   `json:"error"`
	Data     any     `json:"data"`
	Progress float64 `json:"progress"`
}

// BackgroundQueue struct abstracts a background processes queue, including a
// safe map of queue items, and methods to enqueue, dequeue, update and check
// the status of the queue items.
type BackgroundQueue struct {
	processes sync.Map
}

// NewBackgroundQueue function initializes a new queue and return it.
func NewBackgroundQueue() *BackgroundQueue {
	return &BackgroundQueue{}
}

// Enqueue method creates a new queue item and enqueue it into the current
// background queue and returns its ID. It initialize the queue item done to
// false, its error and data to nil. This queue item parameters can be updated
// using the resulting ID and the queue Update method.
func (q *BackgroundQueue) Enqueue() string {
	id := util.RandomHex(QueueIDLen)
	q.processes.Store(id, QueueItem{
		Done:  false,
		Error: nil,
		Data:  make(map[string]any),
	})
	return id
}

// Dequeue method removes a item from que current queue using the id provided.
// It returns if the item was in the queue before remove it.
func (q *BackgroundQueue) Dequeue(id string) bool {
	if _, ok := q.processes.Load(id); !ok {
		return false
	}
	q.processes.Delete(id)
	return true
}

// UpdateProgress method updates the queue item with the provided id with the
// progress provided. It returns if the operation was successful, which means
// that the queue item was in the queue before the update.
func (q *BackgroundQueue) UpdateProgress(id string, progress float64) bool {
	if progress < 0 || progress > 100 {
		return false
	}
	iQueueItem, ok := q.processes.Load(id)
	if !ok {
		return false
	}
	queueItem, ok := iQueueItem.(QueueItem)
	if !ok {
		return false
	}
	if progress > queueItem.Progress {
		queueItem.Progress = progress
		q.processes.Store(id, queueItem)
	}
	return true
}

// UpdateData method updates the queue item with the provided id with the data
// map provided. It returns if the operation was successful, which means that
// the queue item was in the queue before the update.
func (q *BackgroundQueue) UpdateData(id string, data any) bool {
	iQueueItem, ok := q.processes.Load(id)
	if !ok {
		return false
	}
	queueItem, ok := iQueueItem.(QueueItem)
	if !ok {
		return false
	}
	queueItem.Data = data
	q.processes.Store(id, queueItem)
	return true
}

// IsDone method returns if the queue item with the provided id is done or not,
// its progress, data and error. It also returns if the queue item was in the
// queue before the operation.
func (q *BackgroundQueue) IsDone(id string) (QueueItem, bool) {
	iQueueItem, ok := q.processes.Load(id)
	if !ok {
		return QueueItem{}, false
	}
	queueItem, ok := iQueueItem.(QueueItem)
	if !ok {
		return QueueItem{}, false
	}
	return queueItem, true
}

// Fail method updates the queue item with the provided id as failed, and also
// stores the provided error into the queue item. It returns if the operation
// was successful, which means that the queue item was in the queue before the
// update.
func (q *BackgroundQueue) Fail(id string, err error) bool {
	iQueueItem, ok := q.processes.Load(id)
	if !ok {
		return false
	}
	queueItem, ok := iQueueItem.(QueueItem)
	if !ok {
		return false
	}
	queueItem.Error = err
	queueItem.Done = true
	q.processes.Store(id, queueItem)
	return true
}

// Done method updates the queue item with the provided id as done, and also
// stores the provided data map into the queue item. It returns if the operation
// was successful, which means that the queue item was in the queue before the
// update.
func (q *BackgroundQueue) Done(id string, data any) bool {
	iQueueItem, ok := q.processes.Load(id)
	if !ok {
		return false
	}
	queueItem, ok := iQueueItem.(QueueItem)
	if !ok {
		return false
	}
	queueItem.Done = true
	queueItem.Progress = 100
	queueItem.Data = data
	q.processes.Store(id, queueItem)
	return true
}

// StepProgressChannel method returns a channel to update the progress of the
// queue item with the provided id. It returns a channel to send the progress
// updates and it is safe to close it when the progress updates are done.
func (q *BackgroundQueue) StepProgressChannel(id string, step, totalSteps int) chan float64 {
	prgressCh := make(chan float64)
	go func() {
		for progress := range prgressCh {
			qi, ok := q.IsDone(id)
			if !ok || qi.Done {
				close(prgressCh)
				return
			}
			// calc partial progress of current step
			stepIndex := step - 1
			partialStep := 100 / totalSteps
			stepProgress := float64(stepIndex*partialStep) + progress/float64(totalSteps)
			if !q.UpdateProgress(id, float64(stepProgress)) {
				close(prgressCh)
				return
			}
		}
	}()
	return prgressCh
}
