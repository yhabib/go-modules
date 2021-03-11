package pqueue

import (
	"reflect"
	"testing"
)

func TestSize(t *testing.T) {
	pq := New()
	description := "Size on empty Priority Queue"
	expected := 0
	if actual := pq.Size(); actual != expected {
		t.Fatalf("FAIL: %s - actual: %d, expected %d", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestInsert(t *testing.T) {
	pq := New()
	pq.Insert(4)
	pq.Insert(1)
	pq.Insert(9)
	pq.Insert(12)
	pq.Insert(3)
	description := "Add item to Priority Queue"
	expected := []int{0, 1, 3, 9, 12, 4}
	if !reflect.DeepEqual(pq.items, expected) {
		t.Fatalf("FAIL: %s - actual: %d, expected: %d", description, pq.items, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestFindMin(t *testing.T) {
	pq := New()
	pq.Insert(4)
	pq.Insert(1)
	pq.Insert(9)
	pq.Insert(12)
	pq.Insert(3)
	description := "FindMin item in Priority Queue"
	expected := 1
	if actual := pq.FindMin(); actual != expected {
		t.Fatalf("FAIL: %s - actual: %d, expected: %d", description, pq.items, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestDeldMin(t *testing.T) {
	pq := New()
	pq.Insert(4)
	pq.Insert(1)
	pq.Insert(9)
	pq.Insert(12)
	pq.Insert(3)
	description := "FindMin item in Priority Queue"
	expectedValue := 1
	expected := []int{0, 3, 4, 9, 12}
	if actual := pq.DelMin(); actual != expectedValue || !reflect.DeepEqual(pq.items, expected) {
		t.Fatalf("FAIL: %s - actual: %d, expected: %d", description, pq.items, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestIsEmpty(t *testing.T) {
	pq := New()
	description := "IsEmpty"
	expected := true
	if actual := pq.IsEmpty(); !actual {
		t.Fatalf("FAIL: %s - actual: %t, expected: %t", description, actual, expected)
	}
	pq.Insert(4)
	expected = false
	if actual := pq.IsEmpty(); actual {
		t.Fatalf("FAIL: %s - actual: %t, expected: %t", description, actual, expected)
	}
	t.Logf("PASS: %s", description)
}

func TestBuildHeap(t *testing.T) {
	pq := New()
	description := "BuildHeap"
	pq.BuildHeap([]int{9, 6, 5, 2, 3})
	expected := []int{0, 2, 3, 5, 6, 9}

	if !reflect.DeepEqual(pq.items, expected) {
		t.Fatalf("FAIL: %s - actual: %d, expected: %d", description, pq.items, expected)
	}

	t.Logf("PASS: %s", description)
}
