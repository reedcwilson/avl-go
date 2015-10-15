package avl

import (
  "fmt"
  "math"
)

var _ = fmt.Println

// Tree is an AVL Tree
type Tree struct {
  Size int
  root *Node
}

// gets the balance of the given tree
func balance(node *Node) int {
  if node == nil {
    return 0
  }
  return height(node.left) - height(node.right)
}


// Clear sets the tree to nil and resets the size
func (t *Tree) Clear() {
  t.Size = 0
  t.root = nil
}


// Delete removes the element matching the given node or false if it couldn't
// find it
func (t *Tree) Delete(value Comparer) bool {
  var found bool
  t.root, found = del(t.root, &Node{ Value: value })
  if found {
    t.Size--
  }
  return found
}

func del(cur, node *Node) (*Node, bool) {
  var found bool
  // we didn't find it
  if cur == nil {
    return cur, false
  }
  comparison := cur.Value.Compare(node.Value)
  // search on the left if the node is less than the current
  if comparison < 0 {
    cur.left, found = del(cur.left, node)
  // search on the right if the node is greater than the current
  } else if comparison > 0 {
    cur.right, found = del(cur.right, node)
  // otherwise we found it
  } else {
    found = true
    // when we only have one or no branches
    if cur.left == nil || cur.right == nil {
      var temp *Node
      if cur.left == nil {
        temp = cur.right
      } else {
        temp = cur.left
      }
      // there are no children
      if temp == nil {
        temp = cur
        cur = nil
      // there is one child
      } else {
        cur = temp
      }
      temp = nil
    } else {
      temp, _ := min(cur.right)
      cur.Value = temp.Value
      cur.right, _ = del(cur.right, temp)
    }
  }
  // we deleted the last node
  if cur == nil {
    return cur, found
  }
  // update the height
  updateHeight(cur)
  // do the balancing
  bal := balance(cur)
  //fmt.Println(bal, cur)
  rightBal := balance(cur.right)
  leftBal := balance(cur.left)
  // left left case
  if bal > 1 && leftBal >= 0 {
    return rotateRight(cur), found
  }
  // left right case
  if bal > 1 && leftBal < 0 {
    cur.left = rotateLeft(cur.left)
    return rotateRight(cur), found
  }
  // right right case
  if bal < -1 && rightBal <= 0 {
    return rotateLeft(cur), found
  }
  // right left case
  if bal < -1 && rightBal > 0 {
    cur.right = rotateRight(cur.right)
    return rotateLeft(cur), found
  }
  return cur, found
}


// Find returns the element matching the given node or false if it couldn't 
// find it
func (t *Tree) Find(value Comparer) (Comparer, bool) {
  var node *Node
  var ok bool
  node, ok = find(t.root, &Node{ Value: value })
  return node.Value, ok
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
  return node.Height
}

func updateHeight(node *Node) {
  node.Height = int(math.Max(float64(height(node.left)), float64(height(node.right))) + 1)
}


// Insert adds the given node to the tree and performs any necessary rebalancing
func (t *Tree) Insert(value Comparer) bool {
  var added bool
  t.root, added = insert(t.root, &Node{ Value: value })
  if added {
    t.Size++
  }
  return added
}

func insert(cur, node *Node) (*Node, bool) {
  var found bool
  // base case
  if cur == nil {
    return node, true
  }
  comparison := cur.Value.Compare(node.Value)
  // collision (only supports unique values)
  if comparison == 0 {
    return nil, false
  // node is smaller than current
  } else if comparison < 0 {
    cur.left, found = insert(cur.left, node)
  // node is larger than current
  } else {
    cur.right, found = insert(cur.right, node)
  }
  if found {
    updateHeight(cur)
    // do the rebalancing
    bal := balance(cur)
    // left left case
    if bal > 1 && cur.Value.Compare(cur.left.Value) < 0 {
      return rotateRight(cur), found
    }
    // right right case
    if bal < -1 && cur.Value.Compare(cur.right.Value) > 0 {
      return rotateLeft(cur), found
    }
    // left right case
    if bal > 1 && cur.Value.Compare(cur.left.Value) > 0 {
      cur.left = rotateLeft(cur.left)
      return rotateRight(cur), found
    }
    // right left case
    if bal < -1 && cur.Value.Compare(cur.right.Value) < 0 {
      cur.right = rotateRight(cur.right)
      return rotateLeft(cur), found
    }
  }
  return cur, found
}


// Max returns the largest element in the tree or false if the tree is empty
func (t *Tree) Max() (Comparer, bool) {
  var node *Node
  var found bool
  node, found = max(t.root)
  return node.Value, found
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
func (t *Tree) Min() (Comparer, bool) {
  var node *Node
  var ok bool
  node, ok = min(t.root)
  return node.Value, ok
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


func (t *Tree) printTree() {
  queue := Queue{}
  if t.root == nil {
    return
  }
  queue.Push(t.root)
  for queue.Len() > 0 {
    node := queue.Pop()
    fmt.Printf("%v ", node)
    if node.left != nil {
      queue.Push(node.left)
    }
    if node.right != nil {
      queue.Push(node.right)
    }
  }
  fmt.Println()
}


func rotateLeft(node *Node) *Node {
  // store temporary variables
  x := node.right
  y := x.left
  // do the rotation
  x.left = node
  node.right = y
  // update heights
  updateHeight(x)
  updateHeight(y)
  return x
}


func rotateRight(node *Node) *Node {
  // store temporary variables
  x := node.left
  y := x.right
  // do the rotation
  x.right = node
  node.left = y
  // update heights
  updateHeight(x)
  updateHeight(y)
  return x
}
