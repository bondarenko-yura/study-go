package concurrent

import (
	"fmt"
	"math/rand"
	"strings"
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
	r := rand.New(rand.NewSource(777))
	for i := 0; i < 100; i++ {
		size := 1 + r.Intn(10)
		load := size + r.Intn(1000)
		t.Run(fmt.Sprintf("size %d load %d", size, load), func(t *testing.T) {
			c := &Collector{}
			ObjectPool(size, load, c)
			c.ValidateEach(func(v any) {
				if v.(int) < 1 || v.(int) > size {
					t.Errorf("want: 1 <= v <= %d, got: %d", size, v)
				}
			})
		})
	}
}

func TestReadFromClosedChannel(t *testing.T) {
	stream := make(chan int)
	close(stream)
	if v, open := <-stream; open || v != 0 {
		t.Errorf("want: closed channel, got: %v, %v", v, open)
	}
}

func TestWriteToClosedChannel(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic, but no panic occurred.")
		}
	}()
	stream := make(chan int)
	close(stream)
	stream <- 1
}

func TestReadWithTimeout(t *testing.T) {
	c := &Collector{}
	ReadWithTimeout(c)
	c.AssertHasOnlyInAnyOrder(t, "25", "timeout")
}

func TestWriteWithTimeout(t *testing.T) {
	c := &Collector{}
	SelectWithDefault(c)
	c.ValidateEach(func(v any) {
		if !strings.HasPrefix(v.(string), "In default after") {
			t.Errorf("want: In default after, got: %v", v)
		}
	})
}
