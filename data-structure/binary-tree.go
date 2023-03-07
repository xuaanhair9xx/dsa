package main

import (
	"fmt"
)

type BinaryTree struct {
	Data interface {}
	LeftNode *BinaryTree
	RightNode *BinaryTree
}

func create(data interface {}) *BinaryTree {
	return BinaryTree { Data: data, LeftNode: nil, RightNode: nil}
}

func (root *BinaryTree) insertRightNode(data interface {}) {
	node := create(data)
	root.RightNode = node
}

func (root *BinaryTree) insertLeftNode(data interface {}) {
	node := create(data)
	root.LeftNode = node
}
