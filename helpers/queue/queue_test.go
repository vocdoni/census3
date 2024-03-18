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
	_, ok := q.processes.Load(id)
	c.Assert(ok, qt.IsTrue)
	c.Assert(q.Dequeue(id), qt.IsTrue)
	_, ok = q.processes.Load(id)
	c.Assert(ok, qt.IsFalse)
	c.Assert(q.Dequeue(id), qt.IsFalse)
}

func TestDoneIsDone(t *testing.T) {
	c := qt.New(t)

	q := NewBackgroundQueue()

	id := q.Enqueue()
	data := map[string]any{
		"key": "value",
	}
	queueItem, exists := q.IsDone(id)
	c.Assert(exists, qt.IsTrue)
	c.Assert(queueItem.Done, qt.IsFalse)

	c.Assert(q.Done("wrongID", data), qt.IsFalse)
	c.Assert(q.Done(id, data), qt.IsTrue)

	queueItem, exists = q.IsDone(id)
	c.Assert(queueItem.Done, qt.IsTrue)
	c.Assert(queueItem.Progress, qt.Equals, float64(100))
	c.Assert(queueItem.Data, qt.DeepEquals, data)
	c.Assert(queueItem.Error, qt.IsNil)
	c.Assert(exists, qt.IsTrue)

	_, exists = q.IsDone("wrongID")
	c.Assert(exists, qt.IsFalse)

	c.Assert(q.Dequeue(id), qt.IsTrue)
	_, exists = q.IsDone(id)
	c.Assert(exists, qt.IsFalse)
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
					qi, exists := q.IsDone(queueItemId)
					if !exists {
						asyncErrors.Store(queueItemId, fmt.Errorf("expected queue item not found during done check"))
						continue
					}
					// if it is not done, update it to done
					if !qi.Done {
						// if this actions fails create an error
						if !q.Done(queueItemId, nil) {
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
					qi, exists := q.IsDone(queueItemId)
					if !exists {
						asyncErrors.Store(queueItemId, fmt.Errorf("expected queue item not found during done check"))
						continue
					}
					// if it is done, remove it from the queue, and if this action
					// fails, create an error; unless create a nil error
					if qi.Done {
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
