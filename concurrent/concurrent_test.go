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
