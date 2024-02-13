package queue

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

var (
	drDuration  time.Duration
	drConsumers int
)

func init() {
	flag.DurationVar(&drDuration, "queueDataRaceDuration", time.Minute, "queue "+
		"data race test duration (by default test deadline or 1m)")
	flag.IntVar(&drConsumers, "queueDataRaceConsumers", 100, "number of queue "+
		"data race test consumers")
}

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

func TestQueueDataRace(t *testing.T) {
	c := qt.New(t)
	// initialize some variables to store the results of the test and a queue
	var nProcesses int64
	queueItemIdChan := make(chan string)
	q := NewBackgroundQueue()
	// set a context with the test deadline
	deadline := time.Now().Add(drDuration)
	maxDeadline, ok := t.Deadline()
	c.Assert(ok, qt.IsTrue)
	if deadline.Compare(maxDeadline) == 1 {
		deadline = maxDeadline
	}
	// decreaseing by 5 seconds to ensure a gap to check test results
	ctx, cancel := context.WithDeadline(context.Background(), deadline.Add(-5*time.Second))
	defer cancel()
	// launch producers
	producersWg := new(sync.WaitGroup)
	producersWg.Add(1)
	go func() {
		defer producersWg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				atomic.AddInt64(&nProcesses, 1)
				queueItemIdChan <- q.Enqueue()
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	// create and lunch consumers
	var asyncErrors sync.Map
	updatersWg := new(sync.WaitGroup)
	for i := 0; i < drConsumers; i++ {
		updatersWg.Add(1)
		go func() {
			defer updatersWg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case queueItemId, ok := <-queueItemIdChan:
					// wait for a new id
					if !ok {
						return
					}
					// if not exists create an error
					exists, done, data, err := q.Done(queueItemId)
					if !exists {
						asyncErrors.Store(queueItemId, fmt.Errorf("expected queue item not found during done check"))
						continue
					}
					// if it is not done, update it to done
					if !done {
						// if this actions fails create an error
						if !q.Update(queueItemId, true, data, err) {
							asyncErrors.Store(queueItemId, fmt.Errorf("expected queue item not found during update"))
							continue
						}
					}
					// resend it through the channel
					queueItemIdChan <- queueItemId
				}
			}
		}()
	}
	// create and lunch consumers
	dequeuersWg := new(sync.WaitGroup)
	for i := 0; i < drConsumers; i++ {
		dequeuersWg.Add(1)
		go func() {
			defer dequeuersWg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case queueItemId, ok := <-queueItemIdChan:
					// wait for a new id
					if !ok {
						return
					}
					// if not exists create an error
					exists, done, _, _ := q.Done(queueItemId)
					if !exists {
						asyncErrors.Store(queueItemId, fmt.Errorf("expected queue item not found during done check"))
						continue
					}
					// if it is done, remove it from the queue, and if this action
					// fails, create an error; unless create a nil error
					if done {
						if !q.Dequeue(queueItemId) {
							asyncErrors.Store(queueItemId, fmt.Errorf("expected queue item not found during update"))
						} else {
							asyncErrors.Store(queueItemId, nil)
						}
						continue
					}
					// if it is not done, resend it through the channel
					queueItemIdChan <- queueItemId
				}
			}
		}()
	}
	// wait until goroutines finish
	producersWg.Wait()
	updatersWg.Wait()
	dequeuersWg.Wait()
	// check completed processes errors (nil or not)
	completed := []error{}
	asyncErrors.Range(func(key, value any) bool {
		if err, ok := value.(error); ok {
			completed = append(completed, err)
		} else {
			completed = append(completed, nil)
		}
		return true
	})
	// assert number of completed processes
	c.Assert(int64(len(completed)), qt.Equals, nProcesses)
	// assert that every error is nil
	for _, err := range completed {
		c.Assert(err, qt.IsNil)
	}
	t.Logf("Completed with %d processes created!", nProcesses)
}
