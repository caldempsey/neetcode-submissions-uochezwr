type LruNode struct {
	Key   int
	Value int
	Prev  *LruNode
	Next  *LruNode
}

type LRUCache struct {
	Head     *LruNode
	Tail     *LruNode
	Cache    map[int]*LruNode
	Capacity int
}

func Constructor(capacity int) LRUCache {
	head := &LruNode{}
	tail := &LruNode{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Head:     head,
		Tail:     tail,
		Cache:    make(map[int]*LruNode, capacity),
		Capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	n, ok := l.Cache[key]
	if !ok {
		return -1
	}
	l.removeNode(n)
	l.addToBack(n)
	return n.Value
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.Cache[key]; ok {
		n.Value = value
		l.removeNode(n)
		l.addToBack(n)
		return
	}

	if len(l.Cache) >= l.Capacity {
		victim := l.Head.Next
		l.removeNode(victim)
		delete(l.Cache, victim.Key)
	}

	n := &LruNode{Key: key, Value: value}
	l.addToBack(n)
	l.Cache[key] = n
}

func (l *LRUCache) removeNode(n *LruNode) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (l *LRUCache) addToBack(n *LruNode) {
	n.Prev = l.Tail.Prev
	n.Next = l.Tail
	l.Tail.Prev.Next = n
	l.Tail.Prev = n
}
