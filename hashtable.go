package dsa

import "fmt"

type htItem struct {
	key int
	val int
}

func (i *htItem) String() string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%d: %d", i.key, i.val)
}

/**************************************************
 * List specific for Hashtable
 **************************************************/

type htListNode struct {
	val  *htItem
	next *htListNode
}

func (n *htListNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprint(n.val)
}

type htList struct {
	head *htListNode
}

func (l *htList) String() string {
	if l == nil || l.head == nil {
		return "/"
	}
	s := ""
	n := l.head
	for n != nil {
		s += fmt.Sprint(n)
		if n.next != nil {
			s += " -> "
		}
		n = n.next
	}
	return s
}

func (l *htList) InsertEnd(val *htItem) {
	if l == nil {
		return
	}
	newNode := &htListNode{val: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

func (l *htList) DeleteAtPosition(position int) {
	if l == nil || l.head == nil {
		return
	}
	if position == 0 {
		l.head = l.head.next
		return
	}
	node := l.head
	for i := 0; node != nil && i < position-1; i++ {
		node = node.next
	}
	if node == nil || node.next == nil {
		return
	}
	node.next = node.next.next
}

/**************************************************
 * Hashtable
 **************************************************/

type Hashtable struct {
	table    []*htList
	capacity int
	size     int
}

func NewHashtable(capacity int) *Hashtable {
	t := &Hashtable{
		table:    nil,
		capacity: capacity,
		size:     0,
	}
	t.init()
	return t
}

func (t *Hashtable) init() {
	t.capacity = getPrime(t.capacity)
	t.table = make([]*htList, t.capacity, t.capacity)
}

func (t *Hashtable) hashFunction(key int) int {
	return key % t.capacity
}

func (t *Hashtable) String() string {
	var s string
	for i := 0; i < t.capacity; i++ {
		list := t.table[i]
		s += fmt.Sprintf("[%d]: %v\n", i, list)
	}
	return s
}

func (t *Hashtable) Insert(key int, val int) {
	index := t.hashFunction(key)
	if t.table[index] == nil {
		t.table[index] = &htList{}
	}
	node := t.table[index].head
	for node != nil {
		if node.val.key == key {
			node.val.val = val
			return
		}
		node = node.next
	}
	htItem := &htItem{key, val}
	t.table[index].InsertEnd(htItem)
	t.size++
}

func (t *Hashtable) Delete(key int) {
	index := t.hashFunction(key)
	list := t.table[index]
	if list == nil {
		return
	}
	i := 0
	node := list.head
	for node != nil {
		if node.val.key == key {
			break
		}
		node = node.next
		i++
	}
	if node != nil {
		list.DeleteAtPosition(i)
		t.size--
	}
}

func (t *Hashtable) Get(key int) (int, bool) {
	index := t.hashFunction(key)
	list := t.table[index]
	if list == nil {
		return 0, false
	}
	node := list.head
	for node != nil {
		if node.val.key == key {
			return node.val.val, true
		}
		node = node.next
	}
	return 0, false
}

func (t *Hashtable) Size() int {
	return t.size
}
