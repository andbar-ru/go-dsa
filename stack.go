package dsa

type Stack struct {
	items []int
}

// Creating an empty stack
func NewStack() *Stack {
	return new(Stack)
}

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Add element into stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Remove top element from stack
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Return top element of stack without removing it
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Return the size of the stack
func (s *Stack) Size() int {
	return len(s.items)
}
