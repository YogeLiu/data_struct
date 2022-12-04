package test

import (
	"testing"

	"github.com/YogeLiu/data_struct/tree"
)

func TestBST(t *testing.T) {
	arr := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
	bst := tree.BinarySearchTree{}
	for _, item := range arr {
		bst.Insert(item)
	}
	bst.Order()
}
