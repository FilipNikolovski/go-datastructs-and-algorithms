package arraystack_test

import (
	"testing"

	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/stacks/arraystack"
)

func TestArrayStack(t *testing.T) {
	s := arraystack.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	el, err := s.Pop()
	if err != nil {
		t.Errorf("err should be nil, got: %s", err)
		return
	}
	expected := el.(int)
	if expected != 3 {
		t.Errorf("expected 3, got: %d", expected)
		return
	}

	el, err = s.Pop()
	if err != nil {
		t.Errorf("err should be nil, got: %s", err)
		return
	}
	expected = el.(int)
	if expected != 2 {
		t.Errorf("expected 2, got: %d", expected)
		return
	}
	el, err = s.Peek()
	if err != nil {
		t.Errorf("err should be nil, got: %s", err)
		return
	}
	expected = el.(int)
	if expected != 1 {
		t.Errorf("expected 1, got: %d", expected)
		return
	}

	el, err = s.Pop()
	if err != nil {
		t.Errorf("err should be nil, got: %s", err)
		return
	}
	expected = el.(int)
	if expected != 1 {
		t.Errorf("expected 1, got: %d", expected)
		return
	}

	_, err = s.Peek()
	if err == nil {
		t.Errorf("expected error when trying to peek, got nil on empty stack")
		return
	}

	_, err = s.Pop()
	if err == nil {
		t.Errorf("expected error when trying to pop, got nil on empty stack")
		return
	}

	if !s.Empty() {
		t.Error("expected stack to be empty")
		return
	}

	s.Push(1)
	if s.Empty() {
		t.Error("expected stack not to be empty")
		return
	}

	s.Clear()
	if !s.Empty() {
		t.Error("expected stack to be empty")
	}
}
