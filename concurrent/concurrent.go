package concurrent

import (
	"fmt"
	"strconv"
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

func CondSignal(cnt *Counter) {
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

func CondBroadcast(c *Collector) {
	type Button struct {
		clicked *sync.Cond
	}
	button := Button{clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			// wait until this goroutine is running
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.clicked, func() {
		c.Add("Maximizing window")
		clickRegistered.Done()
	})
	subscribe(button.clicked, func() {
		c.Add("Displaying annoying dialog box")
		clickRegistered.Done()
	})
	subscribe(button.clicked, func() {
		c.Add("Mouse clicked")
		clickRegistered.Done()
	})

	button.clicked.Broadcast()
	clickRegistered.Wait()
}

func DoOnce() int {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	return count
}

type dbFactory struct {
	mu  sync.Mutex
	ids int
	dbs []*db
}

func (d *dbFactory) get() *db {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.dbs) == 0 {
		time.Sleep(50 * time.Millisecond)
		d.ids++
		return &db{id: d.ids}
	}
	db := d.dbs[0]
	d.dbs = d.dbs[1:]
	return db
}

func (d *dbFactory) put(db *db) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.dbs = append(d.dbs, db)
}

type db struct {
	id int
}

type dbPool struct {
	cond          *sync.Cond
	size, maxSize int
	pool          *dbFactory
}

func (p *dbPool) get() *db {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for p.size == p.maxSize {
		p.cond.Wait()
	}
	p.size++
	return p.pool.get()
}

func (p *dbPool) Put(d *db) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	p.pool.put(d)
	p.size--
	p.cond.Signal()
}

func newObjectPool(maxSize int) *dbPool {
	return &dbPool{
		maxSize: maxSize,
		pool:    &dbFactory{},
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func ObjectPool(size, load int, c *Collector) {
	op := newObjectPool(size)
	var wg sync.WaitGroup
	for i := 0; i < load; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db := op.get()
			defer op.Put(db)
			c.Add(db.id)
		}()
	}
	wg.Wait()
}

func ReadWithTimeout(c *Collector) {
	ch := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 25
	}()
	for {
		select {
		case v := <-ch:
			c.Add(strconv.Itoa(v))
		case <-time.After(300 * time.Millisecond):
			c.Add("timeout")
			return
		}
	}
}

func SelectWithDefault(c *Collector) {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		c.Add(fmt.Sprintf("In default after %v", time.Since(start)))
	}
}
