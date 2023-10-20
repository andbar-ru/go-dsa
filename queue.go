package dsa

import "fmt"

type Queue struct {
	items []int
}

// Creating an empty queue
func NewQueue() *Queue {
	return new(Queue)
}

// Check if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Add an element into queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Remove and return the first element from queue
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Return the first element of queue without removing it
func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[0], true
}

// Return the size of the queue
func (q *Queue) Size() int {
	return len(q.items)
}

type CircularQueue struct {
	items []int
	front int
	rear  int
}

// Create an empty circular queue of given size
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{items: make([]int, size, size), front: -1, rear: -1}
}

// Check if the circular queue is full
func (q *CircularQueue) IsFull() bool {
	return q.front == q.rear+1 || q.front == 0 && q.rear == cap(q.items)-1
}

// Check if the circular queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.front == -1
}

// Insert an element into the circular queue
func (q *CircularQueue) Enqueue(item int) error {
	if q.IsFull() {
		return fmt.Errorf("circular queue is full")
	}
	if q.front == -1 {
		q.front = 0
	}
	q.rear = (q.rear + 1) % cap(q.items)
	q.items[q.rear] = item
	return nil
}

// Delete and return an element from the circular queue
func (q *CircularQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	item := q.items[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % cap(q.items)
	}
	return item, true
}

// Return the front element of the circular queue without removing it
func (q *CircularQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[q.front], true
}

func (q *CircularQueue) queue() []int {
	items := make([]int, 0, cap(q.items))
	if q.IsEmpty() {
		return items
	}
	for i := q.front; i != q.rear; i = (i + 1) % cap(q.items) {
		items = append(items, q.items[i])
	}
	items = append(items, q.items[q.rear])
	return items
}

type PriorityQueue = Heap

func NewPriorityQueue() *PriorityQueue {
	return new(PriorityQueue)
}

type Deque struct {
	items []int
}

func NewDeque() *Deque {
	return new(Deque)
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) AddRear(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) AddFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) RemoveFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

func (d *Deque) RemoveRear() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

func (d *Deque) Size() int {
	return len(d.items)
}
