package dequeue

import "testing"

func TestSize(t *testing.T) {
	dequeue := Dequeue{}
	description := "Size on empty Dequeue"
	expected := 0
	if actual := dequeue.size(); actual != expected {
		t.Fatalf("FAIL: %s - q.size(): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestAddFront(t *testing.T) {
	dequeue := Dequeue{}
	dequeue.addFront(4)
	dequeue.addFront("hi")
	description := "Add item to front of Dequeue"
	expected := "hi"
	if actual := dequeue; actual != expected {
		t.Fatalf("FAIL: %s - s.enqueue(4): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestDequeue(t *testing.T) {
	dequeue := Dequeue{}
	dequeue.enqueue(4)
	dequeue.dequeue()
	description := "Remove item from front of Dequeue"
	expected := 0
	if actual := dequeue.size(); actual != expected {
		t.Fatalf("FAIL: %s - s.dequeue(): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestIsEmpty(t *testing.T) {
	dequeue := Dequeue{}
	dequeue.addFront(4)
	description := "Check emptyness of queue"
	expected := false
	if actual := dequeue.isEmpty(); actual != expected {
		t.Fatalf("FAIL: %s - s.isEmpty(): %t, expected %t", description, actual, expected)
	}
	dequeue.removeFront()
	expected = true
	if actual := dequeue.isEmpty(); actual != expected {
		t.Fatalf("FAIL: %s - s.isEmpty(): %t, expected %t", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}
