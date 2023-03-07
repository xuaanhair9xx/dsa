package main

import (
	"fmt"
)

type LinkedListSingle struct {
	Data interface{}
	Next *LinkedListSingle
}

func intialSingleLinkedList(data interface{}) *LinkedListSingle{
	return &LinkedListSingle { Data: data, Next: nil}
}

func (node *LinkedListSingle) PrintList() {
	for node != nil {
		fmt.Printf("%v ", node->Data)
		node = node.Next
	}
}

func (node *LinkedListSingle) GetLength() int {
	count := 0
	for node != nil {
		count ++
		node = node.Next
	}

	return count
}

func (node *LinkedListSingle) AccessNodeWithPosition(pos int) *LinkedListSingle {
	var rNode &LinkedListSingle{}
	i := 1
	for i <= pos && node != nil {
		if (pos == i) {
			rNode = node
		}
		node = node.Next
	}
	return rNode
}

type LinkedListDouble struct {
	Data interface{}
	Next *LinkedListDouble
	Prev *LinkedListDouble
}

func intialDoubleLinkedList(data interface {}) {
	return LinkedListDouble {Data: data, Next: nil, Prev: nil}
}

func (node *LinkedListDouble) AddNode(n *LinkedListDouble) {
	node.Next = n
	n.Prev = node
}

func (node *LinkedListDouble) MoveBackPrevNode() {
	node = node.Prev
} 