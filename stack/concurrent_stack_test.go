package stack

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentStack_Push(t *testing.T) {
	const routineCount = 20
	const count = 100
	const exceptedLength = routineCount * count

	var concurrentStack ConcurrentStack[int]
	var waitGroup sync.WaitGroup
	for routineID := 0; routineID < routineCount; routineID++ {
		waitGroup.Add(1)
		go func(base int) {
			base *= count
			for offset := 0; offset < count; offset++ {
				v := base + offset
				fmt.Printf("push %d\n", v)
				concurrentStack.Push(v)
			}
			waitGroup.Done()
		}(routineID)
	}
	waitGroup.Wait()

	if length := concurrentStack.Length(); length != exceptedLength {
		t.Errorf("concurrent stack's length should be %d, not %d\n",
			exceptedLength, length)
	}

	exists := [exceptedLength]bool{}
	for !concurrentStack.IsEmpty() {
		v := concurrentStack.Top()
		concurrentStack.Pop()
		if exists[v] {
			t.Errorf("element %d should only appear once", v)
		}
		exists[v] = true
	}
}

func TestConcurrentStack_SafePop(t *testing.T) {
	const routineCount = 20
	const count = 100
	const initialCapacity = 1234
	concurrentStack := NewConcurrentStack[int](initialCapacity)
	const exceptedLength = routineCount * count
	for i := 0; i < exceptedLength; i++ {
		concurrentStack.Push(i)
	}
	ch := make(chan int)
	for i := 0; i < routineCount; i++ {
		go func(id int) {
			array := [count]int{}
			ok := false
			for offset := 0; offset < count; offset++ {
				array[offset], ok = concurrentStack.SafePop()
				if !ok {
					t.Errorf("concurrent stack should have top element")
				}
			}
			for _, v := range array {
				ch <- v
			}
		}(i)
	}

	exists := [exceptedLength]bool{}
	for i := 0; i < exceptedLength; i++ {
		v := <-ch
		if exists[v] {
			t.Errorf("element %d should only appear once", v)
		}
		exists[v] = true
	}
	for v, e := range exists {
		if !e {
			t.Errorf("element %d should exist", v)
		}
	}
}
