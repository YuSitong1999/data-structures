package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[int]
	t.Logf("stack Length %d Capacity %d\n", stack.Length(), stack.Capacity())
	stack.Push()
	t.Logf("stack Length %d Capacity %d\n", stack.Length(), stack.Capacity())
	stack.Push(1)
	t.Logf("stack Length %d Capacity %d\n", stack.Length(), stack.Capacity())
	stack.Push(2, 3)
	t.Logf("stack Length %d Capacity %d\n", stack.Length(), stack.Capacity())
	stack.Push(4, 5)
	t.Logf("stack Length %d Capacity %d\n", stack.Length(), stack.Capacity())

	for step, i := 1, 5; i > 0; step, i = step+1, i-1 {
		t.Logf("step %d:\n", step)
		if stack.IsEmpty() {
			t.Errorf("stack should not be empty\n")
		}
		if top := stack.Top(); top != i {
			t.Errorf("stack's top element should be %d, not %d\n", i, top)
		}
		stack.Pop()
		t.Logf("stack Length %d Capacity %d\n", stack.Length(), stack.Capacity())
	}

	if !stack.IsEmpty() {
		t.Errorf("stack should be empty\n")
	}
}

func TestStack2(t *testing.T) {
	stack := NewStack[int](1)
	stack.Push()
	stack.Push(1)
	stack.Push(2, 3)
	for step, i := 1, 3; i > 0; step, i = step+1, i-1 {
		t.Logf("step %d:\n", step)
		if stack.IsEmpty() {
			t.Errorf("stack should not be empty\n")
		}
		if top, ok := stack.SafeTop(); !ok {
			t.Errorf("stack's top element should exist\n")
		} else if top != i {
			t.Errorf("stack's top element should be %d, not %d\n", i, top)
		}
		ok := stack.SafePop()
		if !ok {
			t.Errorf("stack pop operation should success\n")
		}
	}

	if _, ok := stack.SafeTop(); ok {
		t.Errorf("stack's top element should not exist\n")
	}
	if ok := stack.SafePop(); ok {
		t.Errorf("stack pop operation should not success\n")
	}
}

func TestStack_Empty(t *testing.T) {
	var stack Stack[int]
	stack.Push(1, 2, 3)
	if length := stack.Length(); length != 3 {
		t.Errorf("stack's Length should be %d, not %d\n", 3, length)
	}
	stack.Empty()
	if length := stack.Length(); length != 0 {
		t.Errorf("stack's Length should be %d, not %d\n", 0, length)
	}
}
