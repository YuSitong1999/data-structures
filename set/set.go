package set

// Set must use NetSet or NetSetWithCapacity to allocate
type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](elements ...T) (s *Set[T]) {
	s = new(Set[T])
	s.data = make(map[T]struct{})
	s.Insert(elements...)
	return s
}

func NewSetWithCapacity[T comparable](capacity int, elements ...T) (s *Set[T]) {
	s = new(Set[T])
	s.data = make(map[T]struct{}, capacity)
	s.Insert(elements...)
	return s
}

func (s *Set[T]) Insert(elements ...T) *Set[T] {
	for _, e := range elements {
		s.data[e] = struct{}{}
	}
	return s
}

func (s *Set[T]) Remove(elements ...T) *Set[T] {
	for _, e := range elements {
		delete(s.data, e)
	}
	return s
}

func (s *Set[T]) Exist(element T) (exist bool) {
	_, exist = s.data[element]
	return
}

func (s *Set[T]) Len() (length int) {
	return len(s.data)
}

func (s *Set[T]) Items() (items []T) {
	for item := range s.data {
		items = append(items, item)
	}
	return
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	return s.Insert(s2.Items()...)
}

func Union[T comparable](s1, s2 *Set[T]) (result *Set[T]) {
	return NewSet[T]().Insert(s1.Items()...).Insert(s2.Items()...)
}

func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	for _, e := range s.Items() {
		if !s2.Exist(e) {
			s.Remove(e)
		}
	}
	return s
}

func Intersection[T comparable](s1, s2 *Set[T]) (result *Set[T]) {
	if s1.Len() > s2.Len() {
		s1, s2 = s2, s1
	}
	result = NewSetWithCapacity[T](s1.Len())
	for _, e := range s1.Items() {
		if s2.Exist(e) {
			result.Insert(e)
		}
	}
	return
}

func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	return s.Remove(s2.Items()...)
}

func Difference[T comparable](s1, s2 *Set[T]) (result *Set[T]) {
	return NewSet[T](s1.Items()...).Remove(s2.Items()...)
}
