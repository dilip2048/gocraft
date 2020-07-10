package LinkedList

import "testing"

func TestLinkedList(t *testing.T) {
	l := &Node{}
	l.InsertNode(1)
	l.InsertNode(2)
	l.InsertNode(3)
	l.InsertNode(4)
	l.PrintList()
	l.PrintHead()
	l = Reverse(l)
	l.PrintList()
}