package stack

import "sync"

type ConcurrentStack[T any] struct {
	sync.RWMutex
	Stack[T]
}

func NewConcurrentStack[T any](capacity int) (s *ConcurrentStack[T]) {
	s = new(ConcurrentStack[T])
	s.data = make([]T, 0, capacity)
	return
}

func (s *ConcurrentStack[T]) readOnlyOperate(f func()) {
	s.RLock()
	defer s.RUnlock()
	f()
}

func (s *ConcurrentStack[T]) readWriteOperate(f func()) {
	s.Lock()
	defer s.Unlock()
	f()
}

func (s *ConcurrentStack[T]) Push(elements ...T) {
	s.readWriteOperate(func() {
		s.Stack.Push(elements...)
	})
}

func (s *ConcurrentStack[T]) Pop() (element T) {
	s.readWriteOperate(func() {
		element = s.Stack.Top()
		s.Stack.Pop()
	})
	return
}

func (s *ConcurrentStack[T]) SafePop() (element T, ok bool) {
	s.readWriteOperate(func() {
		element, ok = s.Stack.SafeTop()
		if !ok {
			return
		}
		s.Stack.SafePop()
	})
	return
}

func (s *ConcurrentStack[T]) IsEmpty() (isEmpty bool) {
	s.readOnlyOperate(func() {
		isEmpty = s.Stack.IsEmpty()
	})
	return
}

func (s *ConcurrentStack[T]) Empty() {
	s.readOnlyOperate(func() {
		s.Stack.Empty()
	})
}

func (s *ConcurrentStack[T]) Length() (length int) {
	s.readOnlyOperate(func() {
		length = s.Stack.Length()
	})
	return
}

func (s *ConcurrentStack[T]) Capacity() (capacity int) {
	s.readOnlyOperate(func() {
		capacity = s.Stack.Capacity()
	})
	return
}
