package dsa

import (
	"fmt"
	"strconv"
)

/**************************************************
 * Heap
 **************************************************/

type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return new(Heap)
}

func (h *Heap) Heapify(i int) {
	size := len(h.arr)
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < size && h.arr[l] > h.arr[largest] {
		largest = l
	}
	if r < size && h.arr[r] > h.arr[largest] {
		largest = r
	}

	if largest != i {
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		h.Heapify(largest)
	}
}

func (h *Heap) Insert(num int) {
	h.arr = append(h.arr, num)
	size := len(h.arr)
	if size > 1 {
		for i := size/2 - 1; i >= 0; i-- {
			h.Heapify(i)
		}
	}
}

func (h *Heap) Delete(num int) {
	size := len(h.arr)
	var i int
	for i = 0; i < size; i++ {
		if h.arr[i] == num {
			break
		}
	}

	h.arr[i], h.arr[size-1] = h.arr[size-1], h.arr[i]
	h.arr = h.arr[:size-1]
	size--

	for i = size/2 - 1; i >= 0; i-- {
		h.Heapify(i)
	}
}

/**************************************************
 * Fibonacci Heap
 **************************************************/

type FibonacciHeapNode struct {
	parent *FibonacciHeapNode
	left   *FibonacciHeapNode
	right  *FibonacciHeapNode
	child  *FibonacciHeapNode
	degree int
	mark   bool
	key    int
}

func NewFibonacciHeapNode(key int) *FibonacciHeapNode {
	node := &FibonacciHeapNode{
		degree: 0,
		mark:   false,
		parent: nil,
		child:  nil,
		key:    key,
	}
	node.left = node
	node.right = node
	return node
}

func (n *FibonacciHeapNode) String() string {
	s := "("
	if n == nil {
		s += ")"
		return s
	}
	curNode := n
	first := true
	for curNode != n || first {
		if first {
			s += "⟷"
			first = false
		}
		s += " " + strconv.Itoa(curNode.key)
		s += fmt.Sprint(curNode.child)
		s += " ⟷"
		curNode = curNode.right
	}
	s += ")"
	return s
}

func (n *FibonacciHeapNode) limitedString(level, maxLevel int) string {
	if level > maxLevel {
		return ""
	}

	s := "("
	if n == nil {
		s += ")"
		return s
	}
	curNode := n
	first := true
	for curNode != n || first {
		if first {
			s += "⟷"
			first = false
		}
		s += " " + strconv.Itoa(curNode.key)
		s += curNode.child.limitedString(level+1, maxLevel)
		s += " ⟷"
		curNode = curNode.right
	}
	s += ")"
	return s
}

type FibonacciHeap struct {
	min   *FibonacciHeapNode
	trace bool
	found *FibonacciHeapNode
}

func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{}
}

func (h *FibonacciHeap) String() string {
	if h == nil || h.min == nil {
		return "()"
	}
	return fmt.Sprint(h.min)
}

func (h *FibonacciHeap) Insert(node *FibonacciHeapNode) {
	if h == nil {
		return
	}
	if h.min == nil {
		h.min = node
		node.left = node
		node.right = node
	} else {
		node.right = h.min
		node.left = h.min.left
		h.min.left.right = node
		h.min.left = node
		if node.key < h.min.key {
			h.min = node
		}
	}
}

func (h *FibonacciHeap) FindMin() (int, bool) {
	if h == nil || h.min == nil {
		return 0, false
	}
	return h.min.key, true
}

func (h *FibonacciHeap) ExtractMin() (int, bool) {
	if h == nil || h.min == nil {
		return 0, false
	}
	min := h.min
	firstChild := min.child
	curChild := firstChild
	if firstChild != nil {
		for ok := true; ok; ok = curChild != nil && curChild != firstChild { // similar to do..while
			right := curChild.right
			curChild.parent = nil
			h.Insert(curChild)
			curChild = right
		}
	}
	min.left.right = min.right
	min.right.left = min.left
	min.child = nil
	if min == min.right {
		h.min = nil
	} else {
		h.min = min.right
		h.consolidate()
	}
	return min.key, true
}

func (h *FibonacciHeap) consolidate() {
	if h == nil || h.min == nil {
		return
	}
	maxDegree := 0
	d2n := make(map[int]*FibonacciHeapNode)
	curNode := h.min
	visited := make(map[*FibonacciHeapNode]struct{})
	for curNode != nil {
		if _, ok := visited[curNode]; ok {
			break
		}
		visited[curNode] = struct{}{}
		x := curNode
		d := x.degree
		for d2n[d] != nil {
			y := d2n[d]
			if x.key > y.key {
				x, y = y, x
				curNode = x
			}
			fibHeapLink(y, x)
			delete(d2n, d)
			d++
		}
		if d > maxDegree {
			maxDegree = d
		}
		d2n[d] = x
		curNode = curNode.right
	}
	h.min = nil
	for i := 0; i <= maxDegree; i++ {
		node := d2n[i]
		if node != nil {
			h.Insert(node)
		}
	}
}

func (h *FibonacciHeap) find(key int) *FibonacciHeapNode {
	if h == nil || h.min == nil {
		return nil
	}
	var found *FibonacciHeapNode

	var findInNode func(key int, node *FibonacciHeapNode)
	findInNode = func(key int, node *FibonacciHeapNode) {
		if node == nil || found != nil {
			return
		}
		curNode := node
		for ok := true; ok; ok = curNode != node && found == nil {
			if curNode.key == key {
				found = curNode
				return
			}
			child := curNode.child
			findInNode(key, child)
			curNode = curNode.right
		}
	}

	findInNode(key, h.min)
	return found
}

func (h *FibonacciHeap) DecreaseKey(key, newKey int) {
	if h == nil || h.min == nil || newKey > key {
		return
	}
	node := h.find(key)
	if node == nil {
		return
	}
	node.key = newKey
	parent := node.parent
	if parent != nil && node.key < parent.key {
		h.cut(node, parent)
		h.cascadingCut(parent)
	}
	if node.key < h.min.key {
		h.min = node
	}
}

func (h *FibonacciHeap) cut(node, parent *FibonacciHeapNode) {
	if h == nil || h.min == nil || node == nil || parent == nil {
		panic("cut: h == nil || h.min == nil || node == nil || parent == nil")
	}
	node.right.left = node.left
	node.left.right = node.right

	parent.degree--
	if parent.child == node {
		if node.right != node {
			parent.child = node.right
		} else {
			parent.child = nil
		}
	}

	node.right = nil
	node.left = nil
	h.Insert(node)
	node.parent = nil
	node.mark = false
}

func (h *FibonacciHeap) cascadingCut(parent *FibonacciHeapNode) {
	if h == nil || h.min == nil || parent == nil {
		panic("cut: h == nil || h.min == nil || parent == nil")
	}
	grandParent := parent.parent
	if grandParent == nil {
		return
	}
	if !parent.mark {
		parent.mark = true
	} else {
		h.cut(parent, grandParent)
		h.cascadingCut(grandParent)
	}
}

func (h *FibonacciHeap) Delete(key int) {
	if h == nil || h.min == nil {
		return
	}
	newKey := h.min.key - 1
	h.DecreaseKey(key, newKey)
	h.ExtractMin()
}

func fibHeapLink(y, x *FibonacciHeapNode) {
	y.left.right = y.right
	y.right.left = y.left

	c := x.child
	if c == nil {
		y.right = y
		y.left = y
	} else {
		y.right = c
		y.left = c.left
		c.left.right = y
		c.left = y
	}
	y.parent = x
	x.child = y
	x.degree = x.degree + 1
	y.mark = false
}

func MergeHeaps(h1, h2 *FibonacciHeap) *FibonacciHeap {
	if h1 == nil || h2 == nil {
		if h1 == nil {
			return h2
		}
		return h1
	}
	merged := NewFibonacciHeap()
	merged.min = h1.min
	if h1.min != nil && h2.min != nil {
		t1 := h1.min.left
		t2 := h2.min.left
		h1.min.left = t2
		t1.right = h2.min
		h2.min.left = t1
		t2.right = h1.min
	}
	if h1.min == nil || (h2.min != nil && h2.min.key < h1.min.key) {
		merged.min = h2.min
	}
	return merged
}
