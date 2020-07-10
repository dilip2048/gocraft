package LinkedList

import "github.com/rs/zerolog/log"

type Node struct {
	next *Node
	data interface{}
}

func (list *Node) insertNode(data interface{}){
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

func (list *Node) printList()  {
	temp := list
	for temp != nil {
		log.Debug().Msgf("%v", temp.data)
		temp = temp.next
	}
}