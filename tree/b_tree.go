package tree

import (
	"fmt"
)

type bst struct {
	value     int
	rightNode *bst
	leftNode  *bst
}

type BinarySearchTree struct {
	root *bst
}

func (tree *BinarySearchTree) Insert(val int) {
	if tree.root == nil {
		tree.root = &bst{value: val}
		return
	}
	if tree.Search(val) {
		return
	}
	ptr := tree.root
	for {
		if val > ptr.value {
			if ptr.rightNode == nil {
				ptr.rightNode = &bst{value: val}
				break
			}
			ptr = ptr.rightNode
		} else {
			if ptr.leftNode == nil {
				ptr.leftNode = &bst{value: val}
				break
			}
			ptr = ptr.leftNode
		}
	}
}

func (tree *BinarySearchTree) Delete(val int) {
	ptr := tree.root
	if ptr == nil {
		return
	}
	parent := tree.root
	for {
		parent = ptr
		if val > ptr.value {
			ptr = ptr.rightNode
		} else {
			ptr = ptr.leftNode
		}
		if ptr == nil {
			return
		}
		if ptr.value == val {
			break
		}
	}
	if ptr.leftNode == nil && ptr.rightNode == nil {
		ptr = nil
		return
	}
	if ptr.leftNode != nil {
		if parent.leftNode.value == ptr.value {
			parent.leftNode = ptr.leftNode
		}
		if parent.rightNode.value == ptr.value {
			parent.rightNode = ptr.leftNode
		}
	}
	if ptr.rightNode != nil {
		if parent.leftNode.value == ptr.value {
			parent.leftNode = ptr.rightNode
		}
		if parent.rightNode.value == ptr.value {
			parent.rightNode = ptr.rightNode
		}
	}
}

func (tree *BinarySearchTree) Search(val int) (exsits bool) {
	ptr := tree.root
	if ptr == nil {
		return
	}
	for {
		if val > ptr.value {
			ptr = ptr.rightNode
		} else {
			ptr = ptr.leftNode
		}
		if ptr == nil {
			return
		}
		if ptr.value == val {
			return true
		}
	}
}

func (tree *BinarySearchTree) Order() {
	s := Stack{}
	s.Push(tree.root)
	ptr := tree.root
	for s.Length() != 0 {
		for ptr.leftNode != nil {
			s.Push(ptr.leftNode)
			ptr = ptr.leftNode
		}
		top := s.Pop().(*bst)
		fmt.Printf("%d ", top.value)
		if top.rightNode != nil {
			ptr = top.rightNode
			s.Push(top.rightNode)
		}
	}
}
