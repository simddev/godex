package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	c := NewCache(5 * time.Second)
	c.Add("key1", []byte("value1"))

	val, ok := c.Get("key1")
	if !ok {
		t.Fatal("expected key1 to be found")
	}
	if string(val) != "value1" {
		t.Errorf("got %q, want %q", val, "value1")
	}

	_, ok = c.Get("missing")
	if ok {
		t.Fatal("expected missing key to not be found")
	}
}

func TestReap(t *testing.T) {
	interval := 50 * time.Millisecond
	c := NewCache(interval)
	c.Add("key1", []byte("value1"))

	time.Sleep(2 * interval)

	_, ok := c.Get("key1")
	if ok {
		t.Fatal("expected key1 to have been reaped")
	}
}
