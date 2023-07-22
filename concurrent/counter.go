package concurrent

import "sync"

type Counter struct {
	rw sync.RWMutex
	v  int
}

func (c *Counter) Inc() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.v++
}

func (c *Counter) Get() int {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.v
}
