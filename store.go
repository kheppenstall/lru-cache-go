package cache

type store map[string]*node

func (s store) delete(k string) {
	delete(s, k)
}
