package linkedlist

import "errors"

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...any) *List {
	list := &List{}
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *List) Push(v any) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	oldHead := l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return oldHead.Value, nil
	}
	l.head = oldHead.next
	l.head.prev = nil
	return oldHead.Value, nil
}

func (l *List) Pop() (any, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	oldTail := l.tail
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return oldTail.Value, nil
	}
	l.tail = oldTail.prev
	l.tail.next = nil
	return oldTail.Value, nil
}

func (l *List) Reverse() {
	curr := l.head
	for curr != nil {
		tmp := curr.prev
		curr.prev = curr.next
		curr.next = tmp
		curr = curr.prev
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Count() int {
	count := 0
	cur := l.head
	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}

func (ll *List) Delete(v any) bool {
	if ll.head == nil {
		return false
	}
	cur := ll.head
	for cur != nil {
		if cur.Value == v {
			if cur.prev == nil {
				ll.head = cur.next
			} else {
				cur.prev.next = cur.next
			}
			if cur.next == nil {
				ll.tail = cur.prev
			} else {
				cur.next.prev = cur.prev
			}
			return true
		}
		cur = cur.next
	}
	return false
}
