package concurrent

import (
	"sync"
	"time"
)

func WaitGroupLoop(c *Collector) {
	var wg sync.WaitGroup
	for _, v := range []string{"a", "b", "c"} {
		wg.Add(1)
		fn := func(in string) {
			defer wg.Done()
			c.Add(in)
		}
		go fn(v)
	}
	wg.Wait()
}

func Generator(from, to int, c *Collector) {
	res := make(chan int, 2)
	c.AddFromChannel(res)
	var wg sync.WaitGroup
	for i := from; i <= to; i++ {
		wg.Add(1)
		go func(in int) {
			defer wg.Done()
			res <- in
		}(i)
	}
	go func() {
		defer close(res)
		wg.Wait()
	}()
}

func Cond(cnt *Counter) {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 2)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		cnt.Inc()
		queue = queue[1:]
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 5; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		queue = append(queue, struct{}{})
		go removeFromQueue(100 * time.Millisecond)
		c.L.Unlock()
	}
}
