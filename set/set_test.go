package set

import "testing"

func TestSet_Insert(t *testing.T) {
	s := NewSet[int]().Insert(1).Insert(2).Insert(3).Insert(2).Insert(1)
	if s.Len() != 3 {
		t.Errorf("the len of the set should be 3")
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet[int](1, 2, 3, 4, 5, 4, 3, 2, 1)
	s.Remove(1)
	s.Remove(2)
	if s.Len() != 3 {
		t.Errorf("the len of the set should be 3")
	}
}

func TestSet_Exist(t *testing.T) {
	s := NewSetWithCapacity(3, 1, 2, 3, 3, 2, 1)
	for i := 1; i <= 3; i++ {
		if !s.Exist(i) {
			t.Errorf("the element %d should exist in the set", i)
		}
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s1.Union(s2)
	if s1.Len() != 5 || !s1.Exist(1) || !s1.Exist(2) ||
		!s1.Exist(3) || !s1.Exist(4) || !s1.Exist(5) {
		t.Errorf("set union error, set: %v\n", s1.Items())
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s1 = Union(s1, s2)
	if s1.Len() != 5 || !s1.Exist(1) || !s1.Exist(2) ||
		!s1.Exist(3) || !s1.Exist(4) || !s1.Exist(5) {
		t.Errorf("set union error, set: %v\n", s1.Items())
	}
}

func TestSet_Intersection(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s1.Intersection(s2)
	if s1.Len() != 1 || s1.Exist(1) || s1.Exist(2) ||
		!s1.Exist(3) || s1.Exist(4) || s1.Exist(5) {
		t.Errorf("set intersection error, set: %v\n", s1.Items())
	}
}
func TestIntersection(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s1 = Intersection(s1, s2)
	if s1.Len() != 1 || s1.Exist(1) || s1.Exist(2) ||
		!s1.Exist(3) || s1.Exist(4) || s1.Exist(5) {
		t.Errorf("set intersection error, set: %v\n", s1.Items())
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s1.Difference(s2)
	if s1.Len() != 2 || !s1.Exist(1) || !s1.Exist(2) ||
		s1.Exist(3) || s1.Exist(4) || s1.Exist(5) {
		t.Errorf("set difference error, set: %v\n", s1.Items())
	}
}

func TestDifference(t *testing.T) {
	s1 := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s1 = Difference(s1, s2)
	if s1.Len() != 2 || !s1.Exist(1) || !s1.Exist(2) ||
		s1.Exist(3) || s1.Exist(4) || s1.Exist(5) {
		t.Errorf("set difference error, set: %v\n", s1.Items())
	}
}
