package LinkedList

import "testing"

func TestLinkedList(t *testing.T) {
	l := &Node{}
	l.insertNode(1)
	l.insertNode(2)
	l.insertNode(3)
	l.insertNode(4)
	l.printList()
}