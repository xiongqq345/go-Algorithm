package _46

type LRUCache struct {
	size, capacity int
	cache          map[int]*LinkedNode
	head, tail     *LinkedNode
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LinkedNode, capacity),
		head:     new(LinkedNode),
		tail:     new(LinkedNode),
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.remove(node)
		c.addToHead(node)
		return node.v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.cache[key]; ok {
		c.remove(n)
	} else {
		c.size++
		if c.size > c.capacity {
			c.removeTail()
			c.size--
		}
	}

	node := new(LinkedNode)
	node.k = key
	node.v = value
	c.addToHead(node)
	c.cache[key] = node
}

type LinkedNode struct {
	k, v       int
	prev, next *LinkedNode
}

func (c *LRUCache) remove(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) removeTail() {
	delete(c.cache, c.tail.prev.k)
	c.remove(c.tail.prev)
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