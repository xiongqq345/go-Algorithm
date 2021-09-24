package jz_offerII

//运用所掌握的数据结构，设计和实现一个  LRU (Least Recently Used，最近最少使用) 缓存机制 。
//
//实现 LRUCache 类：
//
//LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

type LinkedNode struct {
	k, v       int
	prev, next *LinkedNode
}

type LRUCache struct {
	cache      map[int]*LinkedNode
	len, cap   int
	head, tail *LinkedNode
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		cache: make(map[int]*LinkedNode, capacity),
		cap:   capacity,
		head:  new(LinkedNode),
		tail:  new(LinkedNode),
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.cache[key]; ok {
		c.removeNode(n)
		c.addToHead(n)
		return n.v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.cache[key]; ok {
		n.v = value
		c.removeNode(n)
		c.addToHead(n)
		return
	}
	n := &LinkedNode{
		k: key,
		v: value,
	}
	c.addToHead(n)
	c.len++
	c.cache[key] = n
	if c.len > c.cap {
		c.removeTail()
		c.len--
	}
}

func (c *LRUCache) removeTail() {
	delete(c.cache, c.tail.prev.k)
	c.removeNode(c.tail.prev)
}
func (c *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addToHead(node *LinkedNode) {
	node.next = c.head.next
	node.prev = c.head
	node.next.prev = node
	c.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
