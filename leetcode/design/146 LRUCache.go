package design

type LinkedNode struct {
	k, v       int
	prev, next *LinkedNode
}

type LRUCache struct {
	len, cap   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		cap:   capacity,
		cache: make(map[int]*LinkedNode, capacity),
		head:  new(LinkedNode),
		tail:  new(LinkedNode),
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
		c.len++
		if c.len > c.cap {
			c.removeTail()
			c.len--
		}
	}

	node := new(LinkedNode)
	node.k = key
	node.v = value
	c.addToHead(node)
	c.cache[key] = node
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
