package cache

type lruCache struct {
	capacity int
	store    store
	head     *node
	tail     *node
}

func (c *lruCache) put(k string, v int) {
	if c.store[k] != nil {
		c.updateValue(k, v)
		return
	}

	n := c.newNode(k, v)
	c.store[k] = &n

	if c.tail.value == 0 {
		c.setTail(&n)
	}

	c.setHead(&n)

	if len(c.store) > c.capacity {
		c.evict()
	}
}

func (c *lruCache) updateValue(k string, v int) {
	n := c.store[k]
	n.value = v
	c.setHead(n)
}

func (c *lruCache) newNode(k string, v int) node {
	return node{
		next:     &node{},
		previous: c.head,
		key:      k,
		value:    v,
	}
}

func (c *lruCache) setTail(n *node) {
	c.tail = n
	n.previous = &node{}
}

func (c *lruCache) setHead(n *node) {
	if len(c.store) > 1 {
		c.remove(n)
	}
	c.setNodeToHead(n)
}

func (c *lruCache) remove(n *node) {
	if n == c.tail {
		c.setTail(n.next)
	}
	n.previous.next = n.next
	n.next.previous = n.previous
}

func (c *lruCache) setNodeToHead(n *node) {
	n.previous = c.head
	c.head.next = n
	c.head = n
	n.next = &node{}
}

func (c *lruCache) evict() {
	key := c.tail.key
	c.setTail(c.tail.next)
	c.store.delete(key)
}

func (c *lruCache) get(k string) int {
	n := c.store[k]
	if n == nil {
		return -1
	}

	if n == c.tail {
		c.setTail(n.next)
	}

	c.setHead(n)

	return n.value
}
