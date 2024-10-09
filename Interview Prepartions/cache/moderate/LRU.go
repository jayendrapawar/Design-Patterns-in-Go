package moderate

import "fmt"

type Node struct {
    key  int
    val  int
    prev *Node
    next *Node
}

type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node
    tail     *Node
}

func NewLRUCache(cap int) *LRUCache {
    op := &LRUCache{
        capacity: cap,
        cache:    make(map[int]*Node),
        head:     &Node{key: -1, val: -1},
        tail:     &Node{key: -1, val: -1},
    }

    op.head.next = op.tail
    op.tail.prev = op.head
    return op
}

func (op *LRUCache) Get(a int) int {
    if node, exists := op.cache[a]; exists {
        op.moveToFront(node)
        return node.val
    }
    return -1
}

func (op *LRUCache) Put(a, b int) {
    if node, exists := op.cache[a]; exists {
        node.val = b
        op.moveToFront(node)
    } else {
        if len(op.cache) >= op.capacity {
            op.removeLRU()
        }
        newNode := &Node{key: a, val: b}
        op.cache[a] = newNode
        op.addToFront(newNode)
    }
}

// Helper Methods
func (op *LRUCache) addToFront(node *Node) {
    node.prev = op.head
    node.next = op.head.next
    op.head.next.prev = node
    op.head.next = node
}

func (op *LRUCache) moveToFront(node *Node) {
    op.removeNode(node)
    op.addToFront(node)
}

func (op *LRUCache) removeLRU() {
    lruNode := op.tail.prev
    op.removeNode(lruNode)
    delete(op.cache, lruNode.key)
}

func (op *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func ModerateLRUCache() {
    lru := NewLRUCache(2)

    lru.Put(1, 1)
    lru.Put(2, 2)
    fmt.Println(lru.Get(1)) // returns 1

    lru.Put(3, 3)           // evicts key 2
    fmt.Println(lru.Get(2)) // returns -1 (not found)

    lru.Put(4, 4)           // evicts key 1
    fmt.Println(lru.Get(1)) // returns -1 (not found)
    fmt.Println(lru.Get(3)) // returns 3
    fmt.Println(lru.Get(4)) // returns 4
}
