package queue

import (
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestEnqueueDequeue(t *testing.T) {
	c := qt.New(t)

	q := NewBackgroundQueue()

	id := q.Enqueue()
	_, ok := q.processes[id]
	c.Assert(ok, qt.IsTrue)
	c.Assert(q.Dequeue(id), qt.IsTrue)
	_, ok = q.processes[id]
	c.Assert(ok, qt.IsFalse)
	c.Assert(q.Dequeue(id), qt.IsFalse)
}

func TestUpdateDone(t *testing.T) {
	c := qt.New(t)

	q := NewBackgroundQueue()

	id := q.Enqueue()
	exists, done, _, err := q.Done(id)
	c.Assert(exists, qt.IsTrue)
	c.Assert(err, qt.IsNil)
	c.Assert(done, qt.IsFalse)

	c.Assert(q.Update(id, true, nil, fmt.Errorf("test error")), qt.IsTrue)
	exists, done, _, err = q.Done(id)
	c.Assert(exists, qt.IsTrue)
	c.Assert(err, qt.IsNotNil)
	c.Assert(done, qt.IsTrue)

	c.Assert(q.Dequeue(id), qt.IsTrue)
	exists, done, _, err = q.Done(id)
	c.Assert(exists, qt.IsFalse)
	c.Assert(err, qt.IsNil)
	c.Assert(done, qt.IsFalse)
}
