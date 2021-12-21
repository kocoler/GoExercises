package main

// 需要加缓存的，要不然超时
// 可以抽函数出来

type node struct {
	key   int
	value int
	pre   *node
	next  *node
}

type LRUCache struct {
	capacity int
	num      int
	head     *node
	tail     *node
	cache    map[int]*node
}

func Constructor(capacity int) LRUCache {
	head := &node{
		pre:  nil,
		next: nil,
	}
	tail := &node{pre: head, next: nil}
	head.next = tail
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    make(map[int]*node),
	}
}

func (this *LRUCache) Get(key int) int {
	var n *node
	var ok bool

	if n, ok = this.cache[key]; !ok {
		return -1
	}

	n.pre.next = n.next
	n.next.pre = n.pre
	n.next = this.head.next
	this.head.next.pre = n
	this.head.next = n
	n.pre = this.head
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	var n *node
	var ok bool
	if n, ok = this.cache[key]; ok {
		n.value = value

		n.pre.next = n.next
		n.next.pre = n.pre
		n.next = this.head.next
		this.head.next.pre = n
		this.head.next = n
		n.pre = this.head
	} else {
		n = &node{
			key:   key,
			value: value,
			pre:   this.head,
			next:  this.head.next,
		}
		this.head.next.pre = n
		this.head.next = n
		this.cache[key] = n
		this.num++
		if this.num > this.capacity {
			delete(this.cache, this.tail.pre.key)
			this.tail.pre = this.tail.pre.pre
			this.tail.pre.next = this.tail
			this.num--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
