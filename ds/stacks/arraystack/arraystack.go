package arraystack

import "errors"

type ArrayStack struct {
	elements []interface{}
}

func New() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) Push(el interface{}) {
	s.elements = append(s.elements, el)
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack: collection is empty, cannot pop element")
	}
	el := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return el, nil
}

func (s *ArrayStack) Peek() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack: collection is empty, cannot peek element")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *ArrayStack) Empty() bool {
	return len(s.elements) == 0
}

func (s *ArrayStack) Clear() {
	s.elements = nil
}
