package ds

type Queue[T any] struct {
	Slice []T
}

func (q *Queue[T]) Deque() T {
	var itemToPop = q.Slice[0]
	q.Slice = q.Slice[1:len(q.Slice)]

	return itemToPop
}

func (q *Queue[T]) Enque(val T) {
	q.Slice = append(q.Slice, val)
}

func (q *Queue[T]) Len() int {
	return len(q.Slice)
}
