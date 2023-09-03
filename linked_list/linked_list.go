package linkedlist

import "fmt"

type LinkedList[T comparable] interface {
	Append(element T)
	Remove(element T)
}

type LinkedListNode[T comparable] struct {
	element T
	next    *LinkedListNode[T]
}

type LinkedListImpl[T comparable] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

func (list *LinkedListImpl[T]) Add(element T) {
	node := &LinkedListNode[T]{element: element}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
}

func (list *LinkedListImpl[T]) Remove(element T) {
	if list.head.element == element {
		list.head = list.head.next
		return
	}
	for n := list.head; n != nil; n = n.next {
		if n.next.element == element {
			n.next = n.next.next
			return
		}
	}
	// silently does nothing
}

func (list *LinkedListImpl[T]) String() string {
	s := "[ "
	if list.head != nil {
		s += fmt.Sprintf("%+v", list.head.element)
	}
	for n := list.head.next; n != nil; n = n.next {
		s += fmt.Sprintf(", %+v", n.element)
	}
	s += " ]"
	return s
}
