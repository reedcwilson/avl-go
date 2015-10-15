package avl

import (
  "fmt"
  "testing"
)

// BALANCE

func TestBalance(t *testing.T) {
  tree := buildTree()
  rootBalance := balance(tree.root)
  leftBalance := balance(tree.root.left)
  rightBalance := balance(tree.root.right)
  if rootBalance != -1 && leftBalance != 0 && rightBalance != -2 {
    t.Error("Balance is incorrect")
  }
}


// CLEAR

func TestClear(t *testing.T) {
  tree := buildTree()
  tree.Clear()
  if tree.Size != 0 {
    t.Error("Failed to clear the tree")
  }
}


// DELETE

//func TestDeleteRoot(t *testing.T) {
  //tree := buildTree()
  //_ = tree.Delete(testVal(10))
  //if tree.root.Value.Compare(testVal(9)) != 0 {
    //t.Error("Failed to delete root")
  //}
//}

//func TestDeleteLeft(t *testing.T) {
  //tree := buildTree()
  //tree.printTree()
  //fmt.Println()
  //ok := tree.Delete(testVal(7))
  //if ok && tree.root.left.Value.Compare(testVal(3)) != 0 {
    ////fmt.Println(tree.root.left.Value, ok)
    //tree.printTree()
    //t.Error("Failed to delete on left")
  //}
//}

//func TestDeleteRight(t *testing.T) {
  //tree := buildTree()
  //ok := tree.Delete(testVal(15))
  //if ok && tree.root.right.Value.Compare(testVal(12)) != 0 {
    //t.Error("Failed to delete on right")
  //}
//}


// FIND

func TestFindRoot(t *testing.T) {
  tree := buildTree()
  if _, ok := tree.Find(testVal(10)); !ok {
    t.Error("Find failed to locate the root node")
  }
}

func TestFindLeftNode(t *testing.T) {
  tree := buildTree()
  val := 9
  if value, ok := tree.Find(testVal(val)); !ok && value.Compare(testVal(val)) != 0 {
    t.Error("Find failed to locate the node")
  }
}

func TestFindRightNode(t *testing.T) {
  tree := buildTree()
  val := 17
  if value, ok := tree.Find(testVal(val)); !ok && value.Compare(testVal(val)) != 0 {
    t.Error("Find failed to locate the node")
  }
}

func TestNotFound(t *testing.T) {
  tree := buildTree()
  val := 8
  if _, ok := tree.Find(testVal(val)); ok {
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
  tree.Insert(testVal(8))
  if tree.Size != 1 || tree.root.Value.Compare(testVal(8)) != 0 {
    t.Error("Insert root didn't work")
  }
}

func TestInsertLeft(t *testing.T) {
  tree := buildTree()
  added := tree.Insert(testVal(8))
  if !added || tree.root.left.right.left.Value.Compare(testVal(8)) != 0 {
    t.Error("Insert left didn't work")
  }
}

func TestInsertRight(t *testing.T) {
  tree := buildTree()
  added := tree.Insert(testVal(13))
  if !added || tree.root.right.left.right.Value.Compare(testVal(13)) != 0 {
    t.Error("Insert right didn't work")
  }
}


// MAX

func TestMax(t *testing.T) {
  tree := buildTree()
  val := 17
  if value, ok := tree.Max(); !ok && value.Compare(testVal(val)) != 0 {
    t.Error("Could not find the max node")
  }
}


// MIN

func TestMin(t *testing.T) {
  tree := buildTree()
  val := 3
  if value, ok := tree.Min(); !ok && value.Compare(testVal(val)) != 0 {
    t.Error("Could not find the min node")
  }
}


// ROTATE RIGHT

func TestRotateRight(t *testing.T) {
  tree := buildTree()
  node := rotateRight(tree.root)
  if node.Value.Compare(testVal(7)) != 0 {
    t.Error("Rotate right failed")
  }
}


// ROTATE LEFT

func TestRotateLeft(t *testing.T) {
  tree := buildTree()
  node := rotateLeft(tree.root)
  if node.Value.Compare(testVal(15)) != 0 {
    t.Error("Rotate left failed")
  }
}


/*
*   Testing Tree
*
*        10
*      /    \
*     7     15
*    / \   /  \
*   3   9 12   17
*             /
*           16
*
*/
func buildTree() *Tree {
  // create the root
  root := &Node{ Height: 3, Value: testVal(10) }

  // create the left branch
  l1 := &Node{ Height: 1, Value: testVal(7) }
  l2 := &Node{ Height: 0, Value: testVal(3) }
  l3 := &Node{ Height: 0, Value: testVal(9) }
  l1.left = l2
  l1.right = l3
  root.left = l1

  // create the right branch
  r1 := &Node{ Height: 2, Value: testVal(15) }
  r2 := &Node{ Height: 1, Value: testVal(17) }
  r3 := &Node{ Height: 1, Value: testVal(12) }
  r4 := &Node{ Height: 0, Value: testVal(16) }
  r1.right = r2
  r1.left = r3
  r2.left = r4
  root.right = r1

  tree := &Tree{ root: root, Size: 7 }
  return tree
}

// convenience type
type testVal int

// the implementation of Comparer
func (t testVal) Compare(test Comparer) int {
  // assert type (ignore any errors for testing)
  otherTest := test.(testVal)
  if t > otherTest {
    return -1
  } else if t < otherTest {
    return 1
  } else {
    return 0
  }
}

