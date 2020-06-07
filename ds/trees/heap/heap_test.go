package heap

import (
	"fmt"
	"testing"
)

func TestMinIntHeap(t *testing.T) {
	tt := []int{100, 18, 20, 1, 10, 2, 32, 19, 125}
	h := NewFrom(tt)
	fmt.Println(h)

	v := h.Pop()
	if v != 1 {
		t.Errorf("popped value should be 1, got %d", v)
	}
	n := h.Len()
	if n != 8 {
		t.Errorf("expected length of the heap to be 8, got %d", n)
	}
	h.Push(3)

	v = h.Pop()
	if v != 2 {
		t.Errorf("popped value should be 2, got %d", v)
	}
}
