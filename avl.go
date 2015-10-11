package avl

import (
  //"fmt"
)

// Tree is an AVL Tree
type Tree struct {
  Size int
  root *Node
}


// Clear sets the tree to nil and resets the size
func (t *Tree) Clear() {
  t.Size = 0
  t.root = nil
}


// Delete removes the element matching the given node or false if it couldn't
// find it
func (t *Tree) Delete(node *Node) bool {
  return false
}


// Find returns the element matching the given node or false if it couldn't 
// find it
func (t *Tree) Find(node *Node) (*Node, bool) {
  return find(t.root, node)
}

func find(root *Node, node *Node) (*Node, bool) {
  // base case for no matching node
  if root == nil {
    return &Node{}, false
  }
  result := root.Value.Compare(node.Value)
  if result == 0 {
    return root, true
  } else if result < 0 {
    return find(root.left, node)
  } else {
    return find(root.right, node)
  }
}


// Height returns the height of the deepest level in the tree
func (t *Tree) Height() int {
  return height(t.root)
}

func height(node *Node) int {
  if node == nil {
    return -1
  }
  leftHeight := height(node.left)
  rightHeight := height(node.right)
  if leftHeight > rightHeight {
    return leftHeight + 1
  } else {
    return rightHeight + 1
  }
}


// Insert adds the given node to the tree and performs any necessary rebalancing
func (t *Tree) Insert(node *Node) bool {
  // if we don't have a root then put 'er in
  if t.root == nil {
    t.root = node
    t.Size++
    return true
  }
  added := insert(t.root, node)
  if added {
    t.Size++
  }
  return added
}

func insert(cur, node *Node) bool {
  comparison := cur.Value.Compare(node.Value)
  if comparison == 0 {
    return false
  } else if comparison < 0 {
    if cur.left == nil {
      node.parent = cur
      cur.left = node
      return true
    }
    return insert(cur.left, node)
  } else {
    if cur.right == nil {
      node.parent = cur
      cur.right = node
      return true
    }
    return insert(cur.right, node)
  }
}


// Max returns the largest element in the tree or false if the tree is empty
func (t *Tree) Max() (*Node, bool) {
  return max(t.root)
}

func max(node *Node) (*Node, bool) {
  if node == nil {
    return &Node{}, false
  }
  if node.right == nil {
    return node, true
  } else {
    return max(node.right)
  }
}


// Min returns the smallest element in the tree or false if the tree is empty
func (t *Tree) Min() (*Node, bool) {
  return min(t.root)
}

func min(node *Node) (*Node, bool) {
  if node == nil {
    return &Node{}, false
  }
  if node.left == nil {
    return node, true
  } else {
    return max(node.left)
  }
}
