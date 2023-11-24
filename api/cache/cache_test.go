package cache

import (
	"fmt"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func ExampleCache() {
	c := DefaultCache()
	c.Set("key", "value")
	val, ok := c.Get("key")
	fmt.Println(val, ok)
	c.Delete("key")
	fmt.Println(c.Has("key"))
	c.Clear()
	fmt.Println(c.Has("key"))
	// Output:
	// value true
	// false
	// false
}

func TestCache(t *testing.T) {
	c := qt.New(t)
	// create a new cache and set a value
	testCache := DefaultCache()
	defer testCache.Destroy()
	testCache.Set("key", "value")
	// get the value and check if it's correct
	val, ok := testCache.Get("key")
	c.Assert(ok, qt.Equals, true)
	c.Assert(val.(string), qt.Equals, "value")
	// change the value, get it and check if it's correct
	testCache.Set("key", "value2")
	val, ok = testCache.Get("key")
	c.Assert(ok, qt.Equals, true)
	c.Assert(val.(string), qt.Equals, "value2")
	// delete the value and check if it's gone
	testCache.Delete("key")
	c.Assert(testCache.Has("key"), qt.Equals, false)
	// set the value again, clear the cache and check if it's gone
	testCache.Set("key", "value")
	testCache.Clear()
	c.Assert(testCache.Has("key"), qt.Equals, false)
}

func TestCacheViews(t *testing.T) {
	c := qt.New(t)
	// create a new cache with a sanity interval of 1 second and a min views to
	// survive of 1, and defer its destruction
	sanityInterval := time.Second
	testCache := NewCache(sanityInterval, 1)
	defer testCache.Destroy()
	// set a two values and get one of them to increase its views
	testCache.Set("key1", "value")
	testCache.Set("key2", "value")
	_, ok := testCache.Get("key1")
	c.Assert(ok, qt.Equals, true)
	// wait for the sanity to happen
	time.Sleep(sanityInterval)
	// check if the value with more views is still there
	_, ok = testCache.Get("key1")
	c.Assert(ok, qt.Equals, true)
	// check if the value with less views is gone
	_, ok = testCache.Get("key2")
	c.Assert(ok, qt.Equals, false)
	// wait for the sanity to happen again
	time.Sleep(sanityInterval * 5)
	// check if the value with more views is gone now
	_, ok = testCache.Get("key1")
	c.Assert(ok, qt.Equals, false)
}
