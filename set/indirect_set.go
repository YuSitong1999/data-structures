package set

// IndirectSet must use NetIndirectSet or NetIndirectSetWithCapacity to allocate
type IndirectSet[T any, T2 comparable] struct {
	mapping   func(element T) T2
	remapping func(element T2) T
	data      map[T2]struct{}
}

func NewIndirectSet[T any, T2 comparable](mapping func(element T) T2,
	remapping func(representative T2) T, elements ...T) *IndirectSet[T, T2] {
	is := new(IndirectSet[T, T2])
	is.mapping = mapping
	is.remapping = remapping
	is.data = make(map[T2]struct{})
	is.Insert(elements...)
	return is
}

func (is *IndirectSet[T, T2]) Insert(elements ...T) *IndirectSet[T, T2] {
	for _, e := range elements {
		is.data[is.mapping(e)] = struct{}{}
	}
	return is
}

func (is *IndirectSet[T, T2]) InsertRepresentative(representatives ...T2) *IndirectSet[T, T2] {
	for _, r := range representatives {
		is.data[r] = struct{}{}
	}
	return is
}

func (is *IndirectSet[T, T2]) Remove(elements ...T) *IndirectSet[T, T2] {
	for _, e := range elements {
		delete(is.data, is.mapping(e))
	}
	return is
}

func (is *IndirectSet[T, T2]) RemoveRepresentative(representatives ...T2) *IndirectSet[T, T2] {
	for _, r := range representatives {
		delete(is.data, r)
	}
	return is
}

func (is *IndirectSet[T, T2]) Exist(element T) (exist bool) {
	_, exist = is.data[is.mapping(element)]
	return
}

func (is *IndirectSet[T, T2]) ExistRepresentative(representative T2) (exist bool) {
	_, exist = is.data[representative]
	return
}

func (is *IndirectSet[T, T2]) Len() (length int) {
	return len(is.data)
}

func (is *IndirectSet[T, T2]) Items() (items []T) {
	for item := range is.data {
		items = append(items, is.remapping(item))
	}
	return
}

func (is *IndirectSet[T, T2]) ItemsRepresentative() (representatives []T2) {
	for r := range is.data {
		representatives = append(representatives, r)
	}
	return
}

func (is *IndirectSet[T, T2]) Union(is2 *IndirectSet[T, T2]) *IndirectSet[T, T2] {
	return is.InsertRepresentative(is2.ItemsRepresentative()...)
}

func IndirectSetUnion[T any, T2 comparable](is1, is2 *IndirectSet[T, T2]) (result *IndirectSet[T, T2]) {
	return NewIndirectSet[T, T2](is1.mapping, is1.remapping).
		InsertRepresentative(is1.ItemsRepresentative()...).
		InsertRepresentative(is2.ItemsRepresentative()...)
}

func (is *IndirectSet[T, T2]) Intersection(is2 *IndirectSet[T, T2]) *IndirectSet[T, T2] {
	for _, r := range is.ItemsRepresentative() {
		if !is2.ExistRepresentative(r) {
			is.RemoveRepresentative(r)
		}
	}
	return is
}

func IndirectSetIntersection[T any, T2 comparable](is1, is2 *IndirectSet[T, T2]) (result *IndirectSet[T, T2]) {
	if is1.Len() > is2.Len() {
		is1, is2 = is2, is1
	}
	result = NewIndirectSet[T, T2](is1.mapping, is1.remapping)
	for _, r := range is1.ItemsRepresentative() {
		if is2.ExistRepresentative(r) {
			result.InsertRepresentative(r)
		}
	}
	return
}

func (is *IndirectSet[T, T2]) Difference(is2 *IndirectSet[T, T2]) *IndirectSet[T, T2] {
	return is.RemoveRepresentative(is2.ItemsRepresentative()...)
}

func IndirectSetDifference[T any, T2 comparable](is1, is2 *IndirectSet[T, T2]) (result *IndirectSet[T, T2]) {
	return NewIndirectSet[T](is1.mapping, is1.remapping).
		InsertRepresentative(is1.ItemsRepresentative()...).
		RemoveRepresentative(is2.ItemsRepresentative()...)
}
