package dsa

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := NewList(nil)

	// Assign item values
	list.Head = NewListNode(1)
	second := NewListNode(2)
	third := NewListNode(3)

	// Connect nodes
	list.Head.Next = second
	second.Next = third

	listStr := fmt.Sprint(list)
	want := "1 -> 2 -> 3"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}
}

func TestListOperations(t *testing.T) {
	list := NewList(nil)
	list.InsertEnd(1)
	list.InsertFront(2)
	list.InsertFront(3)
	list.InsertEnd(4)
	list.InsertAfter(list.Head.Next, 5)

	listStr := fmt.Sprint(list)
	want := "3 -> 2 -> 5 -> 1 -> 4"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}

	list.DeleteAtPosition(3)
	listStr = fmt.Sprint(list)
	want = "3 -> 2 -> 5 -> 4"
	if listStr != want {
		t.Errorf("list after deleting an element is %s, want %s", listStr, want)
	}

	found := list.Has(3)
	if !found {
		t.Errorf("list.Has(3) returned false, want true")
	}

	list.Sort()
	listStr = fmt.Sprint(list)
	want = "2 -> 3 -> 4 -> 5"
	if listStr != want {
		t.Errorf("list after sorting is %s, want %s", listStr, want)
	}
}

func TestDoublyLinkedList(t *testing.T) {
	// Initialize nodes
	one := NewListNode(1)
	two := NewListNode(2)
	three := NewListNode(3)

	// Connect nodes
	one.Next = two
	two.Next = three
	two.Prev = one
	three.Prev = two

	// Create Doubly Linked List
	list := &DoublyLinkedList{&List{Head: one}}

	listStr := fmt.Sprint(list)
	want := "1 <-> 2 <-> 3"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}
}

func TestDoublyLinkedListOperations(t *testing.T) {
	dLinkedList := NewDoublyLinkedList(nil)

	dLinkedList.InsertEnd(5)
	dLinkedList.InsertFront(1)
	dLinkedList.InsertFront(6)
	dLinkedList.InsertEnd(9)

	dLinkedList.InsertAfter(dLinkedList.Head, 11)
	dLinkedList.InsertAfter(dLinkedList.Head.Next, 15)

	listStr := fmt.Sprint(dLinkedList)
	want := "6 <-> 11 <-> 15 <-> 1 <-> 5 <-> 9"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}

	dLinkedList.DeleteEnd()
	listStr = fmt.Sprint(dLinkedList)
	want = "6 <-> 11 <-> 15 <-> 1 <-> 5"
	if listStr != want {
		t.Errorf("list after deleting the last element is %s, want %s", listStr, want)
	}
}

func TestCircularList(t *testing.T) {
	// Initialize nodes
	one := NewListNode(1)
	two := NewListNode(2)
	three := NewListNode(3)

	// Connect nodes
	one.Next = two
	two.Next = three
	three.Next = one

	// Create Circular List
	list := &CircularList{&List{Head: one}}

	listStr := fmt.Sprint(list)
	want := "-> 1 -> 2 -> 3 ->"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}
}

func TestCircularListOperations(t *testing.T) {
	list := NewCircularList(nil)
	list.InsertFront(6)
	list.InsertEnd(8)
	list.InsertFront(2)
	list.InsertAfter(list.Head, 10)

	listStr := fmt.Sprint(list)
	want := "-> 2 -> 10 -> 6 -> 8 ->"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}

	list.DeleteEnd()
	listStr = fmt.Sprint(list)
	want = "-> 2 -> 10 -> 6 ->"
	if listStr != want {
		t.Errorf("list is %s, want %s", listStr, want)
	}
}
