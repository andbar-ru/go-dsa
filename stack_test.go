package dsa

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	popped, _ := stack.Pop()
	if popped != 4 {
		t.Errorf("stack.Pop() = %d, want %d", popped, 4)
	}
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(stack.items, want) {
		t.Errorf("stack after popping an element = %v, want %v", stack.items, want)
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := new(Stack)
		for n := 0; n < 1000; n++ {
			stack.Push(n)
		}
		for n := 0; n < 1000; n++ {
			stack.Pop()
		}
	}
}
