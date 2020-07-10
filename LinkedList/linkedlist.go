package LinkedList

import "github.com/rs/zerolog/log"

// The node of the linkedlist
type Node struct {
	next *Node
	data interface{}
}


// This will insert node to a linked list
func (list *Node) InsertNode(data interface{}){
	node := &Node{
		next: nil,
		data: data,
	}
	if list.data == nil {
		list.data = node.data
		list.next = node.next
		return
	}
	temp := list
	for temp.next != nil  {
		temp = temp.next
	}
	temp.next = node
}

// This will print the linked list
func (list *Node) PrintList()  {
	temp := list
	for temp != nil {
		log.Debug().Msgf("%v", temp.data)
		temp = temp.next
	}
}

// This will print the head of the linked list
func (list *Node) PrintHead() {
	log.Debug().Msgf("%v", list.data)
}

// This will remove the data from the end
func (list *Node) Remove()  {
	temp := list
	for temp.next.next != nil{
		temp = temp.next
	}
	temp.next = nil
}

// This will reverse the Linked List
func Reverse(list *Node) *Node{
	var prev *Node
	curr := list
	var next *Node
	for curr != nil {
		// before changing the next of current, do store next node
		next = curr.next
		curr.next = prev
		// move previous and current a step ahead
		prev = curr
		curr = next
	}
	return prev
}