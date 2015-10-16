package avl

import (
	"fmt"
)

type Node struct {
	Height int
	Value  Comparer
	tree   Tree
	left   *Node
	right  *Node
}

func (n *Node) String() string {
	var leftVal Comparer = nil
	var rightVal Comparer = nil
	if n.left != nil {
		leftVal = n.left.Value
	}
	if n.right != nil {
		rightVal = n.right.Value
	}
	return fmt.Sprint("{v:", n.Value, ",h:", n.Height, ",l:", leftVal, ",r:", rightVal, "}")
}
