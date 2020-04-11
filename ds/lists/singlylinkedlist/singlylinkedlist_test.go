package singlylinkedlist_test

import (
	"testing"

	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/lists/singlylinkedlist"
)

func TestSinglyLinkedList(t *testing.T) {
	l := singlylinkedlist.New()

	elements := []int{1, 2, 5, 10, 11, 20, 30}

	for _, e := range elements {
		l.Add(e)
	}

	if l.Len() != 7 {
		t.Errorf("expected length to be 7 got %d", l.Len())
	}

	for i, e := range elements {
		expected, err := l.Get(i)
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
			continue
		}

		if expected.(int) != e {
			t.Errorf("expected %d got %d", e, expected)
		}
	}

	_, err := l.Get(-1)
	if err != singlylinkedlist.ErrOutOfBounds {
		t.Errorf("expected ErrOutOfBounds, got %s", err)
	}

	f := l.First()
	if f.Value().(int) != elements[0] {
		t.Errorf("expected first element to be %d, got %d", elements[0], f.Value().(int))
	}

	l.Remove(0)
	l.Remove(l.Len() - 1)

	l.Clear()
	if !l.Empty() {
		t.Errorf("expected list to be empty, got %d", l.Len())
	}

	l.Remove(0)
}
