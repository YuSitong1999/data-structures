package queue

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](capacity int) (q *Queue[T]) {
	q = new(Queue[T])
	q.data = make([]T, 0, capacity)
	return
}

func (q *Queue[T]) Push(elements ...T) {
	q.data = append(q.data, elements...)
}

func (q *Queue[T]) Pop() {
	q.data = q.data[1:]
}

func (q *Queue[T]) Length() (length int) {
	return len(q.data)
}

func (q *Queue[T]) Front() (element T) {
	return q.data[0]
}
