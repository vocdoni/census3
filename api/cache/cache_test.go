package cache

import (
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

func ExampleCache() {
	c := NewCache()
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
	testCache := NewCache()
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
