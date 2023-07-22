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

func TestCond(t *testing.T) {
	c := &Counter{}
	Cond(c)
	if cnt := c.Get(); cnt < 3 || cnt > 5 {
		t.Errorf("want: 3 <= cnt <= 5, got: %d", cnt)
	}
}
