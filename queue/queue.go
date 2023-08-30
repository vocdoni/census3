package queue

import (
	"sync"

	"go.vocdoni.io/dvote/util"
)

const QueueIDLen = 20

type BackgroundQueue struct {
	mtx       *sync.Mutex
	processes map[string]struct {
		done bool
		err  error
	}
}

func NewBackgroundQueue() *BackgroundQueue {
	return &BackgroundQueue{
		mtx: &sync.Mutex{},
		processes: make(map[string]struct {
			done bool
			err  error
		}),
	}
}

func (q *BackgroundQueue) Enqueue() string {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	id := util.RandomHex(QueueIDLen)
	q.processes[id] = struct {
		done bool
		err  error
	}{done: false, err: nil}
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

func (q *BackgroundQueue) Update(id string, done bool, err error) bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if _, ok := q.processes[id]; !ok {
		return false
	}
	q.processes[id] = struct {
		done bool
		err  error
	}{done, err}
	return true
}

func (q *BackgroundQueue) Done(id string) (bool, error, bool) {
	q.mtx.Lock()
	defer q.mtx.Unlock()

	if p, ok := q.processes[id]; ok {
		return p.done, p.err, true
	}
	return false, nil, false
}
