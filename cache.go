package cache

func newCache(c int) lruCache {
	return lruCache{
		capacity: c,
		store:    store{},
		head:     &node{},
		tail:     &node{},
	}
}
