package concurrent

import (
	"testing"
)

func TestWaitGroupLoop(t *testing.T) {
	c := &Collector{}
	WaitGroupLoop(c)
	c.AssertHasOnlyInAnyOrder(t, "a", "b", "c")
}

func TestGenerator(t *testing.T) {
	c := &Collector{}
	Generator(1, 5, c)
	c.AssertHasOnlyInAnyOrder(t, 1, 2, 3, 4, 5)
}

func TestCondSignal(t *testing.T) {
	c := &Counter{}
	CondSignal(c)
	if cnt := c.Get(); cnt < 3 || cnt > 5 {
		t.Errorf("want: 3 <= cnt <= 5, got: %d", cnt)
	}
}

func TestCondBroadcast(t *testing.T) {
	c := &Collector{}
	CondBroadcast(c)
	c.AssertHasOnlyInAnyOrder(t, "Maximizing window", "Displaying annoying dialog box", "Mouse clicked")
}

func TestDoOnce(t *testing.T) {
	if cnt := DoOnce(); cnt != 1 {
		t.Errorf("want: 1, got: %d", cnt)
	}
}

func TestObjectPool(t *testing.T) {
	c := &Collector{}
	ObjectPool(2, c)
	c.ValidateEach(func(v any) {
		if v.(int) != 1 && v.(int) != 2 {
			t.Errorf("want: 1 or 2, got: %d", v.(int))
		}
	})
}
