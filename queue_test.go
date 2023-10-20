package dsa

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(queue.items, want) {
		t.Errorf("queue = %v, want %v", queue.items, want)
	}
	item, _ := queue.Dequeue()
	if item != 1 {
		t.Errorf("queue.Dequeue() = %d, want %d", item, 1)
	}
	want = []int{2, 3, 4, 5}
	if !reflect.DeepEqual(queue.items, want) {
		t.Errorf("stack after removing an element = %v, want %v", queue.items, want)
	}
}

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := new(Queue)
		for n := 0; n < 1000; n++ {
			queue.Enqueue(n)
		}
		for n := 0; n < 1000; n++ {
			queue.Dequeue()
		}
	}
}

func TestCircularQueue(t *testing.T) {
	cqueue := NewCircularQueue(5)

	// Fails because front = -1
	if _, ok := cqueue.Dequeue(); ok {
		t.Errorf("cqueue.Dequeue(): ok = true, want false")
	}

	cqueue.Enqueue(1)
	cqueue.Enqueue(2)
	cqueue.Enqueue(3)
	cqueue.Enqueue(4)
	cqueue.Enqueue(5)

	// Fails to enqueue because front == 0 && rear == CAP - 1
	if err := cqueue.Enqueue(6); err == nil {
		t.Errorf("cqueue.Enqueue(6): err = nil, want \"circular queue is full\"")
	}

	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(cqueue.queue(), want) {
		t.Errorf("cqueue = %v, want %v", cqueue.queue(), want)
	}

	item, _ := cqueue.Dequeue()
	if item != 1 {
		t.Errorf("queue.Dequeue() = %d, want %d", item, 1)
	}
	want = []int{2, 3, 4, 5}
	if !reflect.DeepEqual(cqueue.queue(), want) {
		t.Errorf("cqueue = %v, want %v", cqueue.queue(), want)
	}

	cqueue.Enqueue(7)
	want = []int{2, 3, 4, 5, 7}
	if !reflect.DeepEqual(cqueue.queue(), want) {
		t.Errorf("cqueue = %v, want %v", cqueue.queue(), want)
	}

	// Fails to enqueue because front == rear + 1
	if err := cqueue.Enqueue(8); err == nil {
		t.Errorf("cqueue.Enqueue(8): err = nil, want \"circular queue is full\"")
	}
}

func TestPriorityQueue(t *testing.T) {
	pqueue := NewPriorityQueue()
	pqueue.Insert(3)
	pqueue.Insert(4)
	pqueue.Insert(9)
	pqueue.Insert(5)
	pqueue.Insert(2)

	want := []int{9, 5, 4, 3, 2}
	if !reflect.DeepEqual(pqueue.arr, want) {
		t.Errorf("pqueue = %v, want %v", pqueue.arr, want)
	}

	pqueue.Delete(4)
	want = []int{9, 5, 2, 3}
	if !reflect.DeepEqual(pqueue.arr, want) {
		t.Errorf("pqueue after deleting 4 = %v, want %v", pqueue.arr, want)
	}
}

func TestDeque(t *testing.T) {
	deque := NewDeque()
	if !deque.IsEmpty() {
		t.Errorf("deque.IsEmpty = false, want true")
	}
	deque.AddRear(8)
	deque.AddRear(5)
	deque.AddFront(7)
	deque.AddFront(10)
	if deque.Size() != 4 {
		t.Errorf("deque.Size() = %d, want 4", deque.Size())
	}
	if deque.IsEmpty() {
		t.Errorf("deque.IsEmpty = true, want false")
	}
	deque.AddRear(11)
	item, _ := deque.RemoveRear()
	if item != 11 {
		t.Errorf("deque.RemoveRear() = %d, want 11", item)
	}
	item, _ = deque.RemoveFront()
	if item != 10 {
		t.Errorf("deque.RemoveFront() = %d, want 10", item)
	}
	deque.AddFront(55)
	deque.AddRear(45)
	want := []int{55, 7, 8, 5, 45}
	if !reflect.DeepEqual(deque.items, want) {
		t.Errorf("deque.items = %v, want %v", deque.items, want)
	}
}
