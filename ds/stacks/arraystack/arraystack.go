package arraystack

import "errors"

// ArrayStack is an implementation of a stack.
type ArrayStack struct {
	elements []interface{}
}

// New creates a new array stack.
func New() *ArrayStack {
	return &ArrayStack{}
}

// Push adds a new element to the stack.
func (s *ArrayStack) Push(el interface{}) {
	s.elements = append(s.elements, el)
}

// Pop fetches the top element of the stack and removes it.
func (s *ArrayStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack: collection is empty, cannot pop element")
	}
	el := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return el, nil
}

// Peek returns the top element from the stack, but does not remove it.
func (s *ArrayStack) Peek() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack: collection is empty, cannot peek element")
	}
	return s.elements[len(s.elements)-1], nil
}

// Empty returns true or false whether the stack has zero elements or not.
func (s *ArrayStack) Empty() bool {
	return len(s.elements) == 0
}

// Clear empties the stack.
func (s *ArrayStack) Clear() {
	s.elements = nil
}
