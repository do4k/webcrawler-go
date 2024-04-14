package data

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}

	q.Enqueue("test1")
	if len(q) != 1 {
		t.Errorf("Enqueue failed. Expected length %v, got %v", 1, len(q))
	}

	element, ok := q.Dequeue()
	if !ok || element != "test1" {
		t.Errorf("Dequeue failed. Expected %v, got %v", "test1", element)
	}
	_, ok = q.Dequeue()
	if ok {
		t.Errorf("Dequeue should fail on empty queue")
	}

	q.Enqueue("test2")
	if !q.Contains("test2") {
		t.Errorf("Contains failed. Expected to find %v in queue", "test2")
	}

	if q.Count() != 1 {
		t.Errorf("Count failed. Expected %v, got %v", 1, q.Count())
	}
}
