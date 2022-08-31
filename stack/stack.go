package stack

type Stack[T any] struct {
	data   []T
	length int
}

func NewStack[T any](capacity int) (s *Stack[T]) {
	s = new(Stack[T])
	s.data = make([]T, 0, capacity)
	return
}

func (s *Stack[T]) Push(elements ...T) {
	s.data = append(s.data, elements...)
	s.length += len(elements)
}

func (s *Stack[T]) Pop() {
	s.length--
	s.data = s.data[:s.length]
	return
}

func (s *Stack[T]) SafePop() (ok bool) {
	if ok = s.length != 0; !ok {
		return
	}
	s.length--
	return
}

func (s *Stack[T]) Top() (element T) {
	return s.data[s.length-1]
}

func (s *Stack[T]) SafeTop() (element T, ok bool) {
	if ok = s.length != 0; !ok {
		return
	}
	element = s.data[s.length-1]
	return
}

func (s *Stack[T]) IsEmpty() (isEmpty bool) {
	return s.length == 0
}

func (s *Stack[T]) Empty() {
	s.data = s.data[:0]
}

func (s *Stack[T]) Length() (length int) {
	return len(s.data)
}

func (s *Stack[T]) Capacity() (capacity int) {
	return cap(s.data)
}
