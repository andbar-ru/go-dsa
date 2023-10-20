package dsa

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	tree := NewBinaryTree(1)
	root := tree.Root
	leftNode := root.InsertLeft(12)
	root.InsertRight(9)
	leftNode.InsertLeft(5)
	leftNode.InsertRight(6)

	result := tree.InorderTraversal()
	want := NewList([]int{5, 12, 6, 1, 9})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("tree.InorderTraversal() = %v, want %v", result, want)
	}

	result = tree.PreorderTraversal()
	want = NewList([]int{1, 12, 5, 6, 9})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("tree.PreorderTraversal() = %v, want %v", result, want)
	}

	result = tree.PostorderTraversal()
	want = NewList([]int{5, 6, 12, 9, 1})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("tree.PostorderTraversal() = %v, want %v", result, want)
	}
}

func TestIsFull(t *testing.T) {
	tree := NewBinaryTree(1)
	root := tree.Root
	left := root.InsertLeft(2)
	root.InsertRight(3)

	left.InsertLeft(4)
	leftRight := left.InsertRight(5)
	leftRight.InsertLeft(6)
	leftRight.InsertRight(7)

	isFull := tree.IsFull()
	if !isFull {
		t.Errorf("tree.IsFull() = %t, want %t", isFull, true)
	}
}

func TestIsPerfect(t *testing.T) {
	tree := NewBinaryTree(1)
	root := tree.Root
	left := root.InsertLeft(2)
	right := root.InsertRight(3)
	left.InsertLeft(4)
	left.InsertRight(5)
	right.InsertLeft(6)

	isPerfect := tree.IsPerfect()
	if isPerfect {
		t.Errorf("tree.IsPerfect() = %t, want %t", isPerfect, false)
	}
}

func TestIsComplete(t *testing.T) {
	tree := NewBinaryTree(1)
	root := tree.Root
	left := root.InsertLeft(2)
	right := root.InsertRight(3)
	left.InsertLeft(4)
	left.InsertRight(5)
	right.InsertLeft(6)

	isComplete := tree.IsComplete()
	if !isComplete {
		t.Errorf("tree.IsComplete() = %t, want %t", isComplete, true)
	}
}

func TestIsBalanced(t *testing.T) {
	tree := NewBinaryTree(1)
	root := tree.Root
	left := root.InsertLeft(2)
	root.InsertRight(3)
	left.InsertLeft(4)
	left.InsertRight(5)

	isBalanced := tree.IsBalanced()
	if !isBalanced {
		t.Errorf("tree.IsBalanced() = %t, want %t", isBalanced, true)
	}
}

func TestDeleteBST(t *testing.T) {
	tree := NewBinarySearchTree(8)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(10)
	tree.Insert(14)
	tree.Insert(4)

	inorderTraversal := tree.InorderTraversal()
	want := NewList([]int{1, 3, 4, 6, 7, 8, 10, 14})
	if !reflect.DeepEqual(inorderTraversal, want) {
		t.Errorf("tree.InorderTraversal() = %v, want %v", inorderTraversal, want)
		return
	}
	tree.Delete(10)
	inorderTraversal = tree.InorderTraversal()
	want = NewList([]int{1, 3, 4, 6, 7, 8, 14})
	if !reflect.DeepEqual(inorderTraversal, want) {
		t.Errorf("tree.InorderTraversal() = %v, want %v", inorderTraversal, want)
		return
	}
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree(2)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(8)

	preorderTraversal := tree.PreorderTraversal()
	want := NewList([]int{4, 2, 1, 3, 7, 5, 8})
	if !reflect.DeepEqual(preorderTraversal, want) {
		t.Errorf("tree.PreorderTraversal() = %v, want %v", preorderTraversal, want)
	}
	tree.Delete(3)
	preorderTraversal = tree.PreorderTraversal()
	want = NewList([]int{4, 2, 1, 7, 5, 8})
	if !reflect.DeepEqual(preorderTraversal, want) {
		t.Errorf("tree.PreorderTraversal() = %v, want %v", preorderTraversal, want)
	}
}
