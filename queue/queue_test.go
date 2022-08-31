package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue[int](4)
	const count = 10
	for i := 1; i <= count; i++ {
		queue.Push(i)
		if length := queue.Length(); length != i {
			t.Errorf("queue's length should be %d, not %d\n", i, length)
		}
	}
	for i := 1; i <= count; i++ {
		v := queue.Front()
		if v != i {
			t.Errorf("queue's front element should be %d, not %d\n", i, v)
		}
		queue.Pop()
		if length := queue.Length(); length != count-i {
			t.Errorf("queue's length should be %d, not %d\n", count-i, length)
		}
	}
}
