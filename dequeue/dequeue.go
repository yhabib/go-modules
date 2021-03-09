package dequeue

// Dequeue data structure
type Dequeue struct {
	Dequeue []interface{}
}

func (d *Dequeue) addFront(item interface{}) {
	d.Dequeue = append(d.Dequeue, item)
}

func (d *Dequeue) addRear(item interface{}) {
	tempQueue := []interface{}{item}
	d.Dequeue = append(tempQueue, d.Dequeue...)
}

func (d *Dequeue) removeFront() interface{} {
	length := len(d.Dequeue)
	item := d.Dequeue[length-1]
	d.Dequeue = d.Dequeue[:length-1]
	return item
}

func (d *Dequeue) removeRear() interface{} {
	item := d.Dequeue[0]
	d.Dequeue = d.Dequeue[1:]
	return item
}

func (d *Dequeue) isEmpty() bool {
	return len(d.Dequeue) == 0
}

func (d *Dequeue) size() int {
	return len(d.Dequeue)
}
