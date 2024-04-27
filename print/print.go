package print

import "fmt"

func ToConsole(v ...any) {
	fmt.Println(v)
}

type entry struct {
	next       *entry
	key, value int
}

type MyHashMap struct {
	data  []*entry
	depth int
}

func Constructor() MyHashMap {
	d := make([]*entry, 16)
	for i := 0; i < len(d); i++ {
		d[i] = &entry{key: -1, value: -1}
	}
	return MyHashMap{data: d}
}

func (h *MyHashMap) Put(key int, value int) {
	e := h.head(key)
	d := 0
	for {
		if e.key == key {
			e.value = value
			break
		}
		if e.next == nil {
			e.next = &entry{key: key, value: value}
			break
		}
		d++
		e = e.next
	}
	if d > h.depth {
		h.depth = d
	}
}

func (h *MyHashMap) Get(key int) int {
	e := h.head(key)
	for e != nil {
		if e.key == key {
			return e.value
		}
		e = e.next
	}
	return -1
}

func (h *MyHashMap) Remove(key int) {
	e := h.head(key)
	for e.next != nil {
		if e.next.key == key {
			e.next = e.next.next
			return
		}
		e = e.next
	}
}

func (h *MyHashMap) head(key int) *entry {
	return h.data[key%len(h.data)]
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
