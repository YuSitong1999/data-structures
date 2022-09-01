package set

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Node struct {
	a int
	b int
}

func mapping(node Node) string {
	return fmt.Sprintf("%d:%d", node.a, node.b)
}

func remapping(representative string) (node Node) {
	slice := strings.Split(representative, ":")
	node.a, _ = strconv.Atoi(slice[0])
	node.b, _ = strconv.Atoi(slice[1])
	return
}

func TestIndirectSet_Insert(t *testing.T) {
	s := NewIndirectSet[Node, string](mapping, remapping).
		Insert(Node{1, 2}).Insert(Node{2, 3}).Insert(Node{3, 4}).
		Insert(Node{2, 3}).Insert(Node{1, 2})
	if s.Len() != 3 {
		t.Errorf("the len of the IndirectSet should be 3")
	}
}

func TestIndirectSet_Remove(t *testing.T) {
	s := NewIndirectSet[Node, string](mapping, remapping,
		Node{1, 2}, Node{2, 3}, Node{3, 4}, Node{2, 3}, Node{1, 2})
	s.Remove(Node{1, 2})
	if s.Len() != 2 {
		t.Errorf("the len of the IndirectSet should be 2")
	}
}

func TestIndirectSet_Exist(t *testing.T) {
	s := NewIndirectSet[Node, string](mapping, remapping,
		Node{1, 2}, Node{2, 3}, Node{3, 4}, Node{2, 3}, Node{1, 2})
	for i := 1; i <= 3; i++ {
		if !s.Exist(Node{i, i + 1}) {
			t.Errorf("the element Node{%d, %d} should exist in the IndirectSet", i, i+1)
		}
	}
}

func TestIndirectSet_Union(t *testing.T) {
	s1 := NewIndirectSet[Node, string](mapping, remapping,
		Node{1, 2}, Node{2, 3}, Node{3, 4})
	s2 := NewIndirectSet[Node, string](mapping, remapping,
		Node{3, 4}, Node{4, 5}, Node{5, 6})
	s1.Union(s2)
	if s1.Len() != 5 || !s1.Exist(Node{1, 2}) || !s1.Exist(Node{2, 3}) ||
		!s1.Exist(Node{3, 4}) || !s1.Exist(Node{4, 5}) || !s1.Exist(Node{5, 6}) {
		t.Errorf("IndirectSet union error, set: %v\n", s1.Items())
	}
}

func TestIndirectSetSet_Intersection(t *testing.T) {
	s1 := NewIndirectSet[Node, string](mapping, remapping,
		Node{1, 2}, Node{2, 3}, Node{3, 4})
	s2 := NewIndirectSet[Node, string](mapping, remapping,
		Node{3, 4}, Node{4, 5}, Node{5, 6})
	s1.Intersection(s2)
	if s1.Len() != 1 || s1.Exist(Node{1, 2}) || s1.Exist(Node{2, 3}) ||
		!s1.Exist(Node{3, 4}) || s1.Exist(Node{4, 5}) || s1.Exist(Node{5, 6}) {
		t.Errorf("IndirectSet intersection error, set: %v\n", s1.Items())
	}
}

func TestIndirectSetSet_Difference(t *testing.T) {
	s1 := NewIndirectSet[Node, string](mapping, remapping,
		Node{1, 2}, Node{2, 3}, Node{3, 4})
	s2 := NewIndirectSet[Node, string](mapping, remapping,
		Node{3, 4}, Node{4, 5}, Node{5, 6})
	s1.Difference(s2)
	if s1.Len() != 2 || !s1.Exist(Node{1, 2}) || !s1.Exist(Node{2, 3}) ||
		s1.Exist(Node{3, 4}) || s1.Exist(Node{4, 5}) || s1.Exist(Node{5, 6}) {
		t.Errorf("IndirectSet difference error, set: %v\n", s1.Items())
	}
}
