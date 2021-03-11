package pqueue

// Pqueue implements a priority queue with binary heaps
type Pqueue struct {
	items []int
}

// New returns instance of the primary queue
func New() (pq *Pqueue) {
	return &Pqueue{items: []int{0}}
}

// Size returns the size of the structure
func (pq *Pqueue) Size() int {
	return len(pq.items) - 1
}

// Insert inserts a new item into the queue and keeps the structure
func (pq *Pqueue) Insert(item int) {
	pq.items = append(pq.items, item)
	pq.percolateUp()
}

// FindMin returns the fisrt item in the queue
func (pq *Pqueue) FindMin() int {
	return pq.items[1]
}

// DelMin returns the minimum and removes it from the queue
func (pq *Pqueue) DelMin() int {
	min := pq.FindMin()
	size := pq.Size()
	pq.items[1] = pq.items[size]
	pq.items = pq.items[:size]
	pq.percolateDown(1)
	return min
}

// IsEmpty checks if queueu is empty
func (pq *Pqueue) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *Pqueue) percolateUp() {
	i := pq.Size()

	for int(i/2) > 0 {
		if pq.items[i] < pq.items[int(i/2)] {
			pq.items[i], pq.items[int(i/2)] = pq.items[int(i/2)], pq.items[i]
		}
		i = int(i / 2)
	}
}

// BuildHeap takes a list to intialize a Heap, by calling just Insert -> O(nlogn) but this approach is O(n)
func (pq *Pqueue) BuildHeap(list []int) {
	pq.items = append(pq.items, list...)
	i := len(list)
	for i > 0 {
		pq.percolateDown(i)
		i = i - 1
	}
}

func (pq *Pqueue) findMinChild(i int) int {
	if i*2+1 > pq.Size() {
		return i * 2
	}
	if pq.items[i*2] < pq.items[i*2+1] {
		return i * 2
	}

	return i*2 + 1
}

func (pq *Pqueue) percolateDown(i int) {
	for i*2 <= pq.Size() {
		child := pq.findMinChild(i)
		if pq.items[child] < pq.items[i] {
			pq.items[i], pq.items[child] = pq.items[child], pq.items[i]
		}
		i = child
	}
}
