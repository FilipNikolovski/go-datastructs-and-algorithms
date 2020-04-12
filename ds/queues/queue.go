package queues

import (
	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/lists/singlylinkedlist"
)

type Queue struct {
	list *singlylinkedlist.SinglyLinkedList
}

func New() *Queue {
	return &Queue{list: singlylinkedlist.New()}
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.InsertBeginning(value)
}

func (q *Queue) Dequeue() interface{} {
	if q.list.Empty() {
		return nil
	}

	el := q.list.Last()
	q.list.Remove(q.list.Len() - 1)
	return el.Value()
}

func (q *Queue) Peek() interface{} {
	return q.list.Last().Value()
}

func (q *Queue) Empty() bool {
	return q.list.Empty()
}
