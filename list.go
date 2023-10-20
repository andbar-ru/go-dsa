package dsa

import (
	"fmt"
	"strings"
)

/**************************************************
 * List Node
 **************************************************/

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode // for Doubly Linked List
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprint(n.Val)
}

/**************************************************
 * Singly Linked List
 **************************************************/

type List struct {
	Head *ListNode
}

func NewList(vals []int) *List {
	dummy := new(ListNode)
	curNode := dummy

	for _, num := range vals {
		curNode.Next = &ListNode{Val: num}
		curNode = curNode.Next
	}

	return &List{Head: dummy.Next}
}

func (l *List) String() string {
	if l == nil || l.Head == nil {
		return "nil"
	}

	visited := make(map[*ListNode]struct{})
	var s string
	n := l.Head

	for n != nil {
		s += fmt.Sprint(n)
		if _, ok := visited[n]; ok {
			s += " ...cycled"
			break
		}
		visited[n] = struct{}{}
		if n.Next != nil {
			s += " -> "
		}
		n = n.Next
	}
	return s
}

// Insert a node at the front
func (l *List) InsertFront(val int) {
	if l == nil {
		return
	}
	newNode := NewListNode(val)
	newNode.Next = l.Head
	l.Head = newNode
}

// Insert a node after a specific node
func (l *List) InsertAfter(node *ListNode, val int) {
	if l == nil || node == nil {
		return
	}
	newNode := NewListNode(val)
	newNode.Next = node.Next
	node.Next = newNode
}

// Insert a new node at the end
func (l *List) InsertEnd(val int) {
	if l == nil {
		return
	}
	newNode := NewListNode(val)
	if l.Head == nil {
		l.Head = newNode
		return
	}
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode
}

// Delete the node at the front
func (l *List) DeleteFront() {
	if l == nil || l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

// Delete the node at the end
func (l *List) DeleteEnd() {
	if l == nil || l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	node := l.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	node.Next = nil
}

// Delete an inner node
func (l *List) DeleteAtPosition(position int) {
	if l == nil || l.Head == nil {
		return
	}
	if position == 0 {
		l.DeleteFront()
		return
	}
	node := l.Head
	// Find the node to be deleted
	for i := 0; node != nil && i < position-1; i++ {
		node = node.Next
	}
	// If the node is not present
	if node == nil || node.Next == nil {
		return
	}
	// Remove the node
	node.Next = node.Next.Next
}

// Has the list a node with given value
func (l *List) Has(val int) bool {
	if l == nil {
		return false
	}
	node := l.Head
	for node != nil {
		if node.Val == val {
			return true
		}
		node = node.Next
	}
	return false
}

// Sort the list
func (l *List) Sort() {
	if l == nil || l.Head == nil || l.Head.Next == nil {
		return
	}
	curNode := l.Head
	var indexNode *ListNode

	for curNode != nil {
		// indexNode points to the node next to curNode
		indexNode = curNode.Next

		for indexNode != nil {
			if curNode.Val > indexNode.Val {
				curNode.Val, indexNode.Val = indexNode.Val, curNode.Val
			}
			indexNode = indexNode.Next
		}

		curNode = curNode.Next
	}
}

/**************************************************
 * Doubly Linked List
 **************************************************/

type DoublyLinkedList struct {
	*List
}

func NewDoublyLinkedList(vals []int) *DoublyLinkedList {
	dummy := new(ListNode)
	curNode := dummy

	for i, val := range vals {
		curNode.Next = NewListNode(val)
		if i > 0 {
			curNode.Next.Prev = curNode
		}
		curNode = curNode.Next
	}

	list := &List{Head: dummy.Next}
	return &DoublyLinkedList{list}
}

func (l *DoublyLinkedList) String() string {
	if l == nil || l.Head == nil {
		return "nil"
	}

	visited := make(map[*ListNode]struct{})
	var s string
	n := l.Head

	for n != nil {
		s += fmt.Sprint(n)
		if _, ok := visited[n]; ok {
			s += " ...cycled"
			break
		}
		visited[n] = struct{}{}
		if n.Next != nil {
			s += " <-> "
		}
		n = n.Next
	}
	return s
}

// Insert a node at the front
func (l *DoublyLinkedList) InsertFront(val int) {
	if l == nil {
		return
	}
	newNode := NewListNode(val)
	newNode.Next = l.Head
	if l.Head != nil {
		l.Head.Prev = newNode
	}
	l.Head = newNode
}

// Insert a node after a specific node
func (l *DoublyLinkedList) InsertAfter(node *ListNode, val int) {
	if l == nil || node == nil {
		return
	}
	newNode := NewListNode(val)
	newNode.Next = node.Next
	node.Next = newNode
	newNode.Prev = node
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

// Insert a new node at the end
func (l *DoublyLinkedList) InsertEnd(val int) {
	if l == nil {
		return
	}
	newNode := NewListNode(val)
	if l.Head == nil {
		l.Head = newNode
		return
	}
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode
	newNode.Prev = node
}

// Delete the node at the front
func (l *DoublyLinkedList) DeleteFront() {
	if l == nil || l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	}
}

// Delete an inner node
func (l *DoublyLinkedList) DeleteAtPosition(position int) {
	if l == nil || l.Head == nil {
		return
	}
	if position == 0 {
		l.DeleteFront()
		return
	}
	node := l.Head
	// Find the node to be deleted
	for i := 0; node != nil && i < position; i++ {
		node = node.Next
	}
	// If the node is not present
	if node == nil {
		return
	}
	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}

// Delete the node at the end
func (l *DoublyLinkedList) DeleteEnd() {
	if l == nil || l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Prev.Next = nil
}

/**************************************************
 * Circular Singly Linked List
 **************************************************/

type CircularList struct {
	*List
}

func NewCircularList(vals []int) *CircularList {
	dummy := new(ListNode)
	curNode := dummy

	for _, num := range vals {
		curNode.Next = &ListNode{Val: num}
		curNode = curNode.Next
	}
	curNode.Next = dummy.Next

	list := &List{Head: dummy.Next}
	return &CircularList{list}
}

func (l *CircularList) String() string {
	if l == nil || l.Head == nil {
		return "nil"
	}

	visited := make(map[*ListNode]struct{})
	var s string
	n := l.Head

	for n != nil {
		if _, ok := visited[n]; ok {
			s = "-> " + s
			break
		}
		s += fmt.Sprint(n)
		visited[n] = struct{}{}
		if n.Next != nil {
			s += " -> "
		}
		n = n.Next
	}
	return strings.TrimSpace(s)
}

func (l *CircularList) InsertFront(val int) {
	if l == nil {
		return
	}
	newNode := NewListNode(val)
	if l.Head == nil {
		l.Head = newNode
		l.Head.Next = l.Head
		return
	}
	// Find the last node
	node := l.Head
	for node.Next != l.Head {
		node = node.Next
	}
	node.Next = newNode
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *CircularList) InsertEnd(val int) {
	if l == nil {
		return
	}
	newNode := NewListNode(val)
	if l.Head == nil {
		l.Head = newNode
		l.Head.Next = l.Head
		return
	}
	node := l.Head
	for node.Next != l.Head {
		node = node.Next
	}
	node.Next = newNode
	newNode.Next = l.Head
}

func (l *CircularList) DeleteFront() {
	if l == nil || l.Head == nil {
		return
	}
	// list has single element
	if l.Head.Next == l.Head {
		l.Head = nil
		return
	}
	// Get the last node
	node := l.Head
	for node.Next != l.Head {
		node = node.Next
	}
	l.Head = l.Head.Next
	node.Next = l.Head
}

func (l *CircularList) DeleteAtPosition(position int) {
	if l == nil || l.Head == nil {
		return
	}
	// list has single element
	if l.Head.Next == l.Head {
		l.Head = nil
		return
	}
	node := l.Head
	// Find the node to be deleted
	for i := 0; i < position-1; i++ {
		node = node.Next
	}
	if node.Next == l.Head {
		l.Head = l.Head.Next
	}
	node.Next = node.Next.Next
}

func (l *CircularList) DeleteEnd() {
	if l == nil || l.Head == nil {
		return
	}
	// list has single element
	if l.Head.Next == l.Head {
		l.Head = nil
		return
	}
	// Get the node before last
	node := l.Head
	for node.Next.Next != l.Head {
		node = node.Next
	}
	node.Next = l.Head
}
