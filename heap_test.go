package dsa

import (
	"fmt"
	"reflect"
	"testing"
)

type fHeapNode struct {
	key      int
	children []*fHeapNode
}

func (n *fHeapNode) FibonacciHeapNode() *FibonacciHeapNode {
	var processNode func(node, parent, left, right *FibonacciHeapNode, children []*fHeapNode)
	processNode = func(node, parent, left, right *FibonacciHeapNode, children []*fHeapNode) {
		if parent != nil {
			node.parent = parent
		}
		if left == nil {
			node.left = node
		} else {
			node.left = left
			left.right = node
		}
		if right == nil {
			node.right = node
		} else {
			node.right = right
			right.left = node
		}
		var cLeft, cFirst *FibonacciHeapNode
		for _, child := range children {
			cNode := &FibonacciHeapNode{key: child.key, degree: len(child.children)}
			processNode(cNode, node, cLeft, cFirst, child.children)
			cLeft = cNode
			if cFirst == nil {
				cFirst = cNode
				node.child = cFirst
			}
		}
	}

	node := &FibonacciHeapNode{key: n.key, degree: len(n.children)}
	processNode(node, nil, nil, nil, n.children)
	return node
}

func TestHeap(t *testing.T) {
	heap := NewHeap()
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(9)
	heap.Insert(5)
	heap.Insert(2)

	want := []int{9, 5, 4, 3, 2}
	if !reflect.DeepEqual(heap.arr, want) {
		t.Errorf("heap = %v, want %v", heap.arr, want)
	}

	heap.Delete(4)
	want = []int{9, 5, 2, 3}
	if !reflect.DeepEqual(heap.arr, want) {
		t.Errorf("heap after deleting 4 = %v, want %v", heap.arr, want)
	}
}

func walkFibonacciHeap(fheap *FibonacciHeap, handler func(*FibonacciHeapNode)) {
	if fheap == nil {
		return
	}

	var walkNode func(*FibonacciHeapNode)
	walkNode = func(node *FibonacciHeapNode) {
		if node == nil {
			return
		}
		curNode := node
		checkNode := node
		for ok := true; ok; ok = curNode != checkNode {
			handler(curNode)
			if curNode.child != nil {
				walkNode(curNode.child)
			}
			curNode = curNode.right
		}
	}

	walkNode(fheap.min)
}

func TestFibonacciHeap(t *testing.T) {
	fheap := NewFibonacciHeap()

	res := fmt.Sprint(fheap)
	want := "()"
	if res != want {
		t.Errorf("Empty fheap = %s, want %s", res, want)
	}

	fheap.Insert(NewFibonacciHeapNode(24))
	res = fmt.Sprint(fheap)
	want = "(⟷ 24() ⟷)"
	if res != want {
		t.Errorf("fheap = %s, want %s", res, want)
	}
	min := fheap.min.key
	wantMin := 24
	if min != wantMin {
		t.Errorf("min = %d, want %d", min, wantMin)
	}

	fheap.Insert(NewFibonacciHeapNode(17))
	fheap.Insert(NewFibonacciHeapNode(3))
	fheap.Insert(NewFibonacciHeapNode(23))
	fheap.Insert(NewFibonacciHeapNode(7))
	res = fmt.Sprint(fheap)
	want = "(⟷ 3() ⟷ 17() ⟷ 24() ⟷ 23() ⟷ 7() ⟷)"
	if res != want {
		t.Errorf("fheap = %s, want %s", res, want)
	}
	min = fheap.min.key
	wantMin = 3
	if min != wantMin {
		t.Errorf("min = %d, want %d", min, wantMin)
	}

	// https://cppalgo.blogspot.com/2011/11/fibonacci-heap.html
	fheap = NewFibonacciHeap()
	fheap.Insert(NewFibonacciHeapNode(1))
	fheap.Insert(NewFibonacciHeapNode(6))
	fheap.Insert(NewFibonacciHeapNode(7))
	fheap.Insert(NewFibonacciHeapNode(4))
	fheap.Insert(NewFibonacciHeapNode(2))
	fheap.Insert(NewFibonacciHeapNode(3))
	fheap.Insert(NewFibonacciHeapNode(5))
	res = fmt.Sprint(fheap)
	want = "(⟷ 1() ⟷ 6() ⟷ 7() ⟷ 4() ⟷ 2() ⟷ 3() ⟷ 5() ⟷)"
	if res != want {
		t.Errorf("fheap = %s, want %s", res, want)
	}
	fheap.ExtractMin()
	res = fmt.Sprint(fheap)
	want = "(⟷ 2(⟷ 6(⟷ 7() ⟷) ⟷ 4() ⟷) ⟷ 3(⟷ 5() ⟷) ⟷)"
	if res != want {
		t.Errorf("fheap after ExtractMin = %s, want %s", res, want)
	}
	min = fheap.min.key
	wantMin = 2
	if min != wantMin {
		t.Errorf("min = %d, want %d", min, wantMin)
	}

	k2d := map[int]int{3: 1, 2: 2, 5: 0, 6: 1, 4: 0, 7: 0}
	unvisited := make(map[int]struct{}, len(k2d))
	for key := range k2d {
		unvisited[key] = struct{}{}
	}

	validateDegree := func(node *FibonacciHeapNode) {
		if node == nil {
			return
		}
		key := node.key
		degree, ok := k2d[key]
		if !ok {
			t.Errorf("Unexpected key %d", key)
		}
		if degree != node.degree {
			t.Errorf("For key %d got degree %d, want %d", key, node.degree, degree)
		}
		delete(unvisited, key)
	}

	walkFibonacciHeap(fheap, validateDegree)
	if len(unvisited) > 0 {
		t.Errorf("There are unvisited keys: %v", unvisited)
	}

	fheap = NewFibonacciHeap()

	fhn := &fHeapNode{
		key: 3,
		children: []*fHeapNode{
			{
				key: 18,
				children: []*fHeapNode{
					{
						key: 39,
					},
				},
			},
			{
				key: 52,
			},
			{
				key: 38,
				children: []*fHeapNode{
					{
						key: 41,
					},
				},
			},
		},
	}
	node := fhn.FibonacciHeapNode()
	fheap.Insert(node)

	fhn = &fHeapNode{
		key: 17,
		children: []*fHeapNode{
			{
				key: 30,
			},
		},
	}
	node = fhn.FibonacciHeapNode()
	fheap.Insert(node)

	fhn = &fHeapNode{
		key: 24,
		children: []*fHeapNode{
			{
				key: 26,
				children: []*fHeapNode{
					{
						key: 35,
					},
				},
			},
			{
				key: 46,
			},
		},
	}
	node = fhn.FibonacciHeapNode()
	fheap.Insert(node)

	fheap.Insert(NewFibonacciHeapNode(23))
	fheap.Insert(NewFibonacciHeapNode(7))
	fheap.Insert(NewFibonacciHeapNode(21))

	k2d = map[int]int{23: 0, 7: 0, 21: 0, 3: 3, 18: 1, 39: 0, 52: 0, 38: 1, 41: 0, 17: 1, 30: 0, 24: 2, 26: 1, 35: 0, 46: 0}
	unvisited = make(map[int]struct{}, len(k2d))
	for key := range k2d {
		unvisited[key] = struct{}{}
	}
	walkFibonacciHeap(fheap, validateDegree)
	if len(unvisited) > 0 {
		t.Errorf("There are unvisited keys: %v", unvisited)
	}

	fheap.ExtractMin()
	res = fmt.Sprint(fheap)
	want = "(⟷ 7(⟷ 24(⟷ 26(⟷ 35() ⟷) ⟷ 46() ⟷) ⟷ 17(⟷ 30() ⟷) ⟷ 23() ⟷) ⟷ 18(⟷ 21(⟷ 52() ⟷) ⟷ 39() ⟷) ⟷ 38(⟷ 41() ⟷) ⟷)"
	if res != want {
		t.Errorf("fheap after ExtractMin = %s, want %s", res, want)
	}
	min = fheap.min.key
	wantMin = 7
	if min != wantMin {
		t.Errorf("min = %d, want %d", min, wantMin)
	}
	k2d = map[int]int{7: 3, 24: 2, 26: 1, 35: 0, 46: 0, 17: 1, 30: 0, 23: 0, 18: 2, 21: 1, 52: 0, 39: 0, 38: 1, 41: 0}
	unvisited = make(map[int]struct{}, len(k2d))
	for key := range k2d {
		unvisited[key] = struct{}{}
	}
	walkFibonacciHeap(fheap, validateDegree)
	if len(unvisited) > 0 {
		t.Errorf("There are unvisited keys: %v", unvisited)
	}
}
