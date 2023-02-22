package queue

type Queue []int

func (q Queue) Empty() bool {
	return len(q) == 0

}

// Add new element in Queue
func (q *Queue) Enqueue(v int) {
	(*q) = append((*q), v)
}

// Delete element in queue
func (q *Queue) Dequeue() int {
	v := (*q)[0]
	*q = (*q)[1:len(*q)]
	return v

}
