package cache

type node struct {
	next     *node
	previous *node
	key      string
	value    int
}
