package singlylinkedlist

import (
	"errors"
)

type Element struct {
	next  *Element
	value interface{}
}

type SinglyLinkedList struct {
	first *Element
	last  *Element
	len   int
}

var (
	ErrOutOfBounds = errors.New("index is out of bounds")
)

func New() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Value() interface{} {
	return e.value
}

func (l *SinglyLinkedList) Len() int {
	return l.len
}

func (l *SinglyLinkedList) First() *Element {
	return l.first
}

func (l *SinglyLinkedList) Empty() bool {
	return l.len == 0
}

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

func (l *SinglyLinkedList) Clear() {
	l.len = 0
	l.first = nil
	l.last = nil
}

func outOfRange(index, len int) bool {
	return index < 0 || index >= len
}
