package concurrent

import (
	"testing"
)

func TestWaitGroupLoop(t *testing.T) {
	c := &Collector{}
	waitGroupLoop(c)
	c.AssertHasAllInAnyOrder(t, "a", "b", "c")
}

func TestGenerator(t *testing.T) {
	c := &Collector{}
	generator(1, 5, c)
	c.AssertHasAllInAnyOrder(t, 1, 2, 3, 4, 5)
}
