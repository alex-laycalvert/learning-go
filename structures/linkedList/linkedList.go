package linkedlist

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

type LinkedList[T any] struct {
	head *listNode[T]
}

func (list *LinkedList[T]) Add(val T) {
	if list.head == nil {
		list.head = &listNode[T]{
			val:  val,
			next: nil,
		}
	}
	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &listNode[T]{
		val:  val,
		next: nil,
	}
}

func (list *LinkedList[T]) Get(index int) (T, bool) {
	var noop T
	if index < 0 || list.head == nil {
		return noop, false
	}
	curr := list.head
	for i := 0; i <= index; i++ {
		curr = curr.next
		if curr == nil {
			return noop, false
		}
	}
	return curr.val, true
}

func (list *LinkedList[T]) Length() int {
	length := 0
	for curr := list.head; curr != nil; curr = curr.next {
	}
	return length
}

func (list *LinkedList[T]) Iter() chan T {
	ch := make(chan T)
	go func() {
		for curr := list.head; curr != nil; curr = curr.next {
			ch <- curr.val
		}
		close(ch)
	}()
	return ch
}
