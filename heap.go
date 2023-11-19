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
			h.Insert(curChild)
			curChild.parent = nil
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
	checkNode := h.min
	for ok := true; ok; ok = curNode != nil && curNode != checkNode {
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
			if len(d2n) == 0 {
				checkNode = x
			}
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
