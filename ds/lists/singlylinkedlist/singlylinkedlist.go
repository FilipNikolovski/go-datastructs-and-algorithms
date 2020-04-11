package singlylinkedlist

import (
	"errors"
)

// Element is an element of a singly linked list.
type Element struct {
	next  *Element
	value interface{}
}

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
	first *Element
	last  *Element
	len   int
}

var (
	// ErrOutOfBounds is an error that is returned if an index is out of bounds.
	ErrOutOfBounds = errors.New("list: index is out of bounds")
)

// New creates a new list.
func New() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

// Next returns the next element.
func (e *Element) Next() *Element {
	return e.next
}

// Value returns the element's value.
func (e *Element) Value() interface{} {
	return e.value
}

// Len returns the length of the list.
func (l *SinglyLinkedList) Len() int {
	return l.len
}

// First returns the first element of the list.
func (l *SinglyLinkedList) First() *Element {
	return l.first
}

// Empty returns true or false whether the list has zero elements or not.
func (l *SinglyLinkedList) Empty() bool {
	return l.len == 0
}

// Add adds a new value to the end of the list and increments the length.
func (l *SinglyLinkedList) Add(value interface{}) {
	newEl := &Element{value: value}

	if l.len == 0 {
		l.first = newEl
	} else {
		l.last.next = newEl
	}

	l.last = newEl
	l.len++
}

// Get fetches the element of the list by the given index.
// If the index is out of range, ErrOutOfBounds is returned.
func (l *SinglyLinkedList) Get(index int) (interface{}, error) {
	if outOfRange(index, l.Len()) {
		return nil, ErrOutOfBounds
	}

	e := l.first
	for i := 0; i < index; i++ {
		e = e.Next()
	}

	return e.Value(), nil
}

// Remove removes an element from the list, by the given index.
func (l *SinglyLinkedList) Remove(index int) {
	if outOfRange(index, l.Len()) {
		return
	}

	var beforeEl *Element
	e := l.first
	for i := 0; i < index; i++ {
		beforeEl = e
		e = e.Next()
	}

	if e == l.first {
		l.first = e.next
	}
	if e == l.last {
		l.last = beforeEl
	}
	if beforeEl != nil {
		beforeEl.next = e.next
	}

	e = nil
	l.len--
}

// Clear empties the list and sets the length to 0.
func (l *SinglyLinkedList) Clear() {
	l.len = 0
	l.first = nil
	l.last = nil
}

func outOfRange(index, len int) bool {
	return index < 0 || index >= len
}
