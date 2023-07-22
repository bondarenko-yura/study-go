package concurrent

import (
	"reflect"
	"sync"
	"testing"
)

type Collector struct {
	rw   sync.RWMutex
	data []interface{}
}

func (c *Collector) Add(v any) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.data = append(c.data, v)
}

func (c *Collector) AddFromChannel(ch chan int) {
	c.rw.Lock()
	go func() {
		defer c.rw.Unlock()
		for i := range ch {
			c.data = append(c.data, i)
		}
	}()
}

func (c *Collector) AssertHasOnlyInAnyOrder(t *testing.T, want ...any) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	if len(c.data) != len(want) {
		t.Errorf("length mismatch, want: %d, got %d", len(want), len(c.data))
	}
	used := make([]bool, len(c.data))
	for _, e := range want {
		found := false
		for i, d := range c.data {
			if !used[i] && reflect.DeepEqual(d, e) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			t.Errorf("want in any order: %#v, got: %#v", want, c.data)
		}
	}
}

func (c *Collector) ValidateEach(validate func(v any)) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	for _, v := range c.data {
		validate(v)
	}
}
