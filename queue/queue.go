package queue

import (
	"sync"

	"go.vocdoni.io/dvote/util"
)

const QueueIDLen = 20

type QueueItem struct {
	done bool
	err  error
	data map[string]any
}

type BackgroundQueue struct {
	mtx       *sync.Mutex
	processes map[string]QueueItem
}

func NewBackgroundQueue() *BackgroundQueue {
	return &BackgroundQueue{
		mtx:       &sync.Mutex{},
		processes: make(map[string]QueueItem),
	}
}

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

func (q *BackgroundQueue) Dequeue(id string) bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if _, ok := q.processes[id]; !ok {
		return false
	}
	delete(q.processes, id)
	return true
}

func (q *BackgroundQueue) Update(id string, done bool, data map[string]any, err error) bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if _, ok := q.processes[id]; !ok {
		return false
	}
	q.processes[id] = QueueItem{done: done, err: err, data: data}
	return true
}

func (q *BackgroundQueue) Done(id string) (bool, bool, map[string]any, error) {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if p, ok := q.processes[id]; ok {
		return true, p.done, p.data, p.err
	}
	return false, false, nil, nil
}
