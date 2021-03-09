package queue

import "testing"

func TestSize(t *testing.T) {
	queue := Queue{}
	description := "Size on empty Queue"
	expected := 0
	if actual := queue.Size(); actual != expected {
		t.Fatalf("FAIL: %s - q.size(): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestEnqueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(4)
	description := "Add item to Queue"
	expected := 1
	if actual := queue.Size(); actual != expected {
		t.Fatalf("FAIL: %s - s.enqueue(4): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestDequeue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(4)
	queue.Dequeue()
	description := "Remove item from Queue"
	expected := 0
	if actual := queue.Size(); actual != expected {
		t.Fatalf("FAIL: %s - s.dequeue(): %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestIsEmpty(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(4)
	description := "Check emptyness of queue"
	expected := false
	if actual := queue.IsEmpty(); actual != expected {
		t.Fatalf("FAIL: %s - s.isEmpty(): %t, expected %t", description, actual, expected)
	}
	queue.Dequeue()
	expected = true
	if actual := queue.IsEmpty(); actual != expected {
		t.Fatalf("FAIL: %s - s.isEmpty(): %t, expected %t", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}
