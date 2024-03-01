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
	done bool
	err  error
	data map[string]any
}

// BackgroundQueue struct abstracts a background processes queue, including a
// mutex to make the operations over it safely and also a map of items,
// identified by the queue item id's.
type BackgroundQueue struct {
	mtx       *sync.Mutex
	processes map[string]QueueItem
}

// NewBackgroundQueue function initializes a new queue and return it.
func NewBackgroundQueue() *BackgroundQueue {
	return &BackgroundQueue{
		mtx:       &sync.Mutex{},
		processes: make(map[string]QueueItem),
	}
}

// Enqueue method creates a new queue item and enqueue it into the current
// background queue and returns its ID. It initialize the queue item done to
// false, its error and data to nil. This queue item parameters can be updated
// using the resulting ID and the queue Update method.
func (q *BackgroundQueue) Enqueue() string {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	id := util.RandomHex(QueueIDLen)
	q.processes[id] = QueueItem{
		done: false,
		err:  nil,
		data: make(map[string]any),
	}
	return id
}

// Dequeue method removes a item from que current queue using the id provided.
// It returns if the item was in the queue before remove it.
func (q *BackgroundQueue) Dequeue(id string) bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if _, ok := q.processes[id]; !ok {
		return false
	}
	delete(q.processes, id)
	return true
}

// Update method updates the information of a queue item identified by the
// provided id. It changes the done, data and error parameters of the found
// queue item to the provided values.
func (q *BackgroundQueue) Update(id string, done bool, data map[string]any, err error) bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if _, ok := q.processes[id]; !ok {
		return false
	}
	q.processes[id] = QueueItem{done: done, err: err, data: data}
	return true
}

// Done method returns the queue item information such as it is done or not, if
// it throws any error and its data. But all this information is only returned
// if the queue item exists in the queue, returned as first parameter.
func (q *BackgroundQueue) Done(id string) (bool, bool, map[string]any, error) {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if p, ok := q.processes[id]; ok {
		return true, p.done, p.data, p.err
	}
	return false, false, nil, nil
}
