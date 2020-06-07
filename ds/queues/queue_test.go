package queues_test

import (
	"testing"

	"github.com/FilipNikolovski/go-datastructs-and-algorithms/ds/queues"
)

func TestQueue(t *testing.T) {
	q := queues.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Empty() {
		t.Error("expected queue not to be empty")
	}

	v := q.Peek()
	if v.(int) != 1 {
		t.Errorf("expected peek value to be 1, got %d", v.(int))
	}

	v = q.Dequeue()
	if v.(int) != 1 {
		t.Errorf("expected dequeued value to be 1, got %d", v.(int))
	}

	v = q.Dequeue()
	if v.(int) != 2 {
		t.Errorf("expected dequeued value to be 2, got %d", v.(int))
	}

	v = q.Dequeue()
	if v.(int) != 3 {
		t.Errorf("expected dequeued value to be 3, got %d", v.(int))
	}

	v = q.Dequeue()
	if v != nil {
		t.Errorf("expected dequeued value to be nil, got %v", v)
	}

	v = q.Peek()
	if v != nil {
		t.Errorf("expected peek value to be nil, got %v", v)
	}

	if !q.Empty() {
		t.Error("expected queue to be empty")
	}

}
