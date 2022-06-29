package solutions

import (
	"container/list"
)

type valueObject struct {
	key   int
	value int
}

type LRUCache struct {
	cap    int
	linked *list.List
	mapped map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:    capacity,
		linked: list.New(),
		mapped: make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	el, ok := this.mapped[key]
	if !ok {
		return -1
	}
	this.linked.MoveToFront(el)

	val := el.Value.(valueObject).value
	return val
}

func (this *LRUCache) Put(key int, value int) {
	if el, ok := this.mapped[key]; ok {
		el.Value = valueObject{key, value}
		this.linked.MoveToFront(el)
		return
	}

	listval := valueObject{key: key, value: value}
	if this.linked.Len() < this.cap {
		el := this.linked.PushFront(listval)
		this.mapped[key] = el
		return
	}

	lastEl := this.linked.Back()
	lastVal := lastEl.Value.(valueObject)
	delete(this.mapped, lastVal.key)
	this.linked.Remove(lastEl)
	el := this.linked.PushFront(listval)
	this.mapped[key] = el
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
