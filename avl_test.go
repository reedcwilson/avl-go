package avl

import (
  "fmt"
  "testing"
)


// CLEAR

func TestClear(t *testing.T) {
  tree := buildTree()
  tree.Clear()
  if tree.Size != 0 {
    t.Error("Failed to clear the tree")
  }
}


// DELETE

func TestDelete(t *testing.T) {
  if false {
    t.Error("Failure")
  }
}


// FIND

func TestFindRoot(t *testing.T) {
  tree := buildTree()
  if _, ok := tree.Find(&Node{ Value: testVal(10) }); !ok {
    t.Error("Find failed to locate the root node")
  }
}

func TestFindLeftNode(t *testing.T) {
  tree := buildTree()
  val := 9
  if node, ok := tree.Find(&Node{ Value: testVal(val) }); !ok && node.Value.Compare(testVal(val)) != 0 {
    t.Error("Find failed to locate the node")
  }
}

func TestFindRightNode(t *testing.T) {
  tree := buildTree()
  val := 17
  if node, ok := tree.Find(&Node{ Value: testVal(val) }); !ok && node.Value.Compare(testVal(val)) != 0 {
    t.Error("Find failed to locate the node")
  }
}

func TestNotFound(t *testing.T) {
  tree := buildTree()
  val := 8
  if _, ok := tree.Find(&Node{ Value: testVal(val) }); ok {
    t.Error("Found a node when it shouldn't have")
  }
}


// HEIGHT

func TestHeight(t *testing.T) {
  tree := buildTree()
  if height := tree.Height(); height != 3 {
    t.Error(fmt.Printf("Failed to find the correct height. Expected: %d, Found: %d", 3, height))
  }
}


// INSERT

func TestInsertRoot(t *testing.T)  {
  tree := new(Tree)
  if tree.Size != 0 {
    t.Error("The tree wasn't initialized correctly")
  }
  tree.Insert(&Node{ Value: testVal(8) })
  if tree.Size != 1 || tree.root.Value.Compare(testVal(8)) != 0 {
    t.Error("Insert didn't work")
  }
}

func TestInsertLeft(t *testing.T) {
  tree := buildTree()
  node := &Node{ Value: testVal(8) }
  added := tree.Insert(node)
  if !added || node.parent.Value.Compare(testVal(9)) != 0 {
    t.Error("Insert didn't work")
  }
}

func TestInsertRight(t *testing.T) {
  tree := buildTree()
  node := &Node{ Value: testVal(13) }
  added := tree.Insert(node)
  if !added || node.parent.Value.Compare(testVal(15)) != 0 {
    t.Error("Insert didn't work")
  }
}


// MAX

func TestMax(t *testing.T) {
  tree := buildTree()
  val := 17
  if node, ok := tree.Max(); !ok && node.Value.Compare(testVal(val)) != 0 {
    t.Error("Could not find the max node")
  }
}


// MIN

func TestMin(t *testing.T) {
  tree := buildTree()
  val := 3
  if node, ok := tree.Min(); !ok && node.Value.Compare(testVal(val)) != 0 {
    t.Error("Could not find the min node")
  }
}


/*
*   Testing Tree
*
*        10
*      /    \
*     7     15
*    / \      \
*   3   9      17
*             /
*           16
*
*/
func buildTree() *Tree {
  // create the root
  root := &Node{ Value: testVal(10) }

  // create the left branch
  l1 := &Node{ Value: testVal(7), parent: root }
  l2 := &Node{ Value: testVal(3), parent: l1 }
  l3 := &Node{ Value: testVal(9), parent: l1 }
  l1.left = l2
  l1.right = l3
  root.left = l1

  // create the right branch
  r1 := &Node{ Value: testVal(15), parent: root }
  r2 := &Node{ Value: testVal(17), parent: r1 }
  r3 := &Node{ Value: testVal(16), parent: r2 }
  r1.right = r2
  r2.left = r3
  root.right = r1

  tree := &Tree{ root: root, Size: 7 }
  return tree
}

// convenience type
type testVal int

// the implementation of Comparer
func (t testVal) Compare(test Comparer) int {
  // assert type (ignore any errors for testing)
  otherTest, _ := test.(testVal)
  if t > otherTest {
    return -1
  } else if t < otherTest {
    return 1
  } else {
    return 0
  }
}

