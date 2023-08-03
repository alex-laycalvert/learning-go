package queue

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.queue = append(q.queue, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *Queue[T]) Top() T {
	var top T
	if len(q.queue) > 0 {
		top = q.queue[0]
	}
	return top
}

func (q *Queue[T]) At(i int) T {
	var item T
	if len(q.queue) <= i {
		return item
	}
	item = q.queue[i]
	return item
}

func (q *Queue[T]) Len() int {
	return len(q.queue)
}

func (q *Queue[T]) Iter() chan T {
	ch := make(chan T)
	go func() {
		for i := 0; i < len(q.queue); i++ {
			ch <- q.queue[i]
		}
		close(ch)
	}()
	return ch
}
