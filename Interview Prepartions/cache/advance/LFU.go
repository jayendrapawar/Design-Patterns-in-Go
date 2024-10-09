package advance

import "fmt"

type Node struct {
	key  int
	val  int
	freq int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head: head, tail: tail}
}

func (dll *DoublyLinkedList) InsertHead(node *Node) {
	node.prev = dll.head
	node.next = dll.head.next
	dll.head.next.prev = node
	dll.head.next = node
}

func (dll *DoublyLinkedList) RemoveNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (dll *DoublyLinkedList) RemoveTail() int {
	node := dll.tail.prev
	dll.RemoveNode(node)
	return node.key
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head.next == dll.tail
}

type LFUCache struct {
	freqMap    map[int]*DoublyLinkedList
	nodeMap    map[int]*Node
	capacity   int
	curSize    int
	leastFreq  int
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		freqMap:   make(map[int]*DoublyLinkedList),
		nodeMap:   make(map[int]*Node),
		capacity:  capacity,
		leastFreq: 0,
	}
}

func (cache *LFUCache) Get(key int) int {
	node, exists := cache.nodeMap[key]
	if !exists {
		return -1
	}

	dll := cache.freqMap[node.freq]
	dll.RemoveNode(node)

	if node.freq == cache.leastFreq && dll.IsEmpty() {
		cache.leastFreq++
	}

	node.freq++
	if _, exists := cache.freqMap[node.freq]; !exists {
		cache.freqMap[node.freq] = NewDoublyLinkedList()
	}

	cache.freqMap[node.freq].InsertHead(node)
	return node.val
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}

	node, exists := cache.nodeMap[key]
	if !exists {
		cache.curSize++
		if cache.curSize > cache.capacity {
			tailKey := cache.freqMap[cache.leastFreq].RemoveTail()
			delete(cache.nodeMap, tailKey)
			cache.curSize--
		}

		newNode := &Node{key: key, val: value, freq: 1}
		if _, exists := cache.freqMap[1]; !exists {
			cache.freqMap[1] = NewDoublyLinkedList()
		}

		cache.freqMap[1].InsertHead(newNode)
		cache.nodeMap[key] = newNode
		cache.leastFreq = 1
	} else {
		node.val = value
		dll := cache.freqMap[node.freq]
		dll.RemoveNode(node)

		if node.freq == cache.leastFreq && dll.IsEmpty() {
			cache.leastFreq++
		}

		node.freq++
		if _, exists := cache.freqMap[node.freq]; !exists {
			cache.freqMap[node.freq] = NewDoublyLinkedList()
		}

		cache.freqMap[node.freq].InsertHead(node)
	}
}

func AdvanceLFUCache() {
	cache := NewLFUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)          // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	cache.Put(4, 4)          // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
