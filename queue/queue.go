// Package queue implements the Queue ADS with two Stacks
package queue

import "github.com/yhabib/go-modules/stack"

// Queue data structure
type Queue struct {
	rear  stack.Stack
	front stack.Stack
}

// Enqueue adds item to the rear of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.rear.Push(item)
}

// Dequeue removes and returns item from the front of the queue
func (q *Queue) Dequeue() interface{} {
	if q.front.IsEmpty() {
		for q.rear.Size() > 0 {
			q.front.Push(q.rear.Pop())
		}
	}
	return q.front.Pop()
}

// IsEmpty chceks the emptyness of the queue
func (q *Queue) IsEmpty() bool {
	return q.rear.IsEmpty() && q.front.IsEmpty()
}

// Size returns size of the Queue
func (q *Queue) Size() int {
	return q.rear.Size() + q.front.Size()
}
