type LruNode struct {
	Key   int
	Value int
	Prev  *LruNode
	Next  *LruNode
}

func NewLruNode(k, v int) *LruNode {
	return &LruNode{
		Key:   k,
		Value: v,
	}
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

	// Touching this node — move it to the back (most recently used)
	l.cacheNode(n)
	return n.Value
}

func (l *LRUCache) Put(key int, value int) {
	n, ok := l.Cache[key]
	if ok {
		// Key exists — update value and move to back
		n.Value = value
		l.cacheNode(n)
		return
	}

	// Key doesn't exist — evict first if at capacity, then insert
	l.evictLeastUsed()
	n = NewLruNode(key, value)
	l.pushNodeToBack(n)
	l.Cache[key] = n
}

func (l *LRUCache) cacheNode(n *LruNode) {
	l.unlinkNode(n)
	l.pushNodeToBack(n)
}

func (l *LRUCache) pushNodeToBack(n *LruNode) {
	n.Prev = l.Tail.Prev
	n.Next = l.Tail
	l.Tail.Prev.Next = n
	l.Tail.Prev = n
}

func (l *LRUCache) unlinkNode(n *LruNode) {
	if n.Prev == nil || n.Next == nil {
		return
	}
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (l *LRUCache) evictLeastUsed() {
	if len(l.Cache) < l.Capacity {
		return
	}
	victim := l.Head.Next
	l.unlinkNode(victim)
	delete(l.Cache, victim.Key)
}