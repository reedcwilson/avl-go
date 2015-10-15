package avl

import (
	"fmt"
)

type Node struct {
  Height int
  Value Comparer
  tree Tree
  left *Node
  right *Node
}

func (n *Node) String() string {
  return fmt.Sprint("{v:", n.Value, ",h:", n.Height, "}")
}

