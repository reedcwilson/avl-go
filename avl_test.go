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

func TestDeleteRoot(t *testing.T) {
	tree := buildTree()
	ok := tree.Delete(testVal(10))
	if ok && tree.root.Value.Compare(testVal(12)) != 0 {
		t.Error("Failed to delete root")
	}
}

func TestDeleteLeft(t *testing.T) {
	tree := buildTree()
	ok := tree.Delete(testVal(7))
	if ok && tree.root.left.Value.Compare(testVal(9)) != 0 {
		t.Error("Failed to delete on left")
	}
}

func TestDeleteRight(t *testing.T) {
	tree := buildTree()
	ok := tree.Delete(testVal(16))
	if ok && tree.root.right.right.left != nil {
		t.Error("Failed to delete on right")
	}
}

func TestDeleteRightLeftCase(t *testing.T) {
	tree := buildTree()
	ok := tree.Delete(testVal(15))
	if ok && tree.root.right.Value.Compare(testVal(16)) != 0 {
		t.Error("Failed the right left delete case")
	}
}

func TestDeleteRightRightCase(t *testing.T) {
	tree := buildTree()
	// prune the children of 7
	tree.root.left.left = nil
	tree.root.left.right = nil
	tree.root.left.Height = 0
	// add a right child to 17
	tree.root.right.right.right = &Node{Value: testVal(20)}
	ok := tree.Delete(testVal(20))
	if ok && tree.root.Value.Compare(testVal(15)) != 0 {
		tree.printTree()
		t.Error("Failed the right right delete case")
	}
}

func TestDeleteLeftRightCase(t *testing.T) {
	tree := buildTree()
	// make 15 a leaf
	tree.root.right.left = nil
	tree.root.right.right = nil
	tree.root.right.Height = 0
	// add two children to 3
	tree.root.left.left.left = &Node{Value: testVal(1)}
	tree.root.left.left.right = &Node{Value: testVal(5)}
	tree.root.left.left.Height = 1
	tree.root.left.Height = 2
	ok := tree.Delete(testVal(15))
	if ok && tree.root.Value.Compare(testVal(7)) != 0 {
		t.Error("Failed the left right delete case")
	}
}

func TestDeleteLeftLeftCase(t *testing.T) {
	tree := buildTree()
	// make 15 a leaf
	tree.root.right.left = nil
	tree.root.right.right = nil
	tree.root.right.Height = 0
	// add two children to 9
	tree.root.left.right.left = &Node{Value: testVal(8.5)}
	tree.root.left.right.right = &Node{Value: testVal(9.5)}
	tree.root.left.right.Height = 1
	tree.root.left.Height = 2
	ok := tree.Delete(testVal(3))
	if ok && tree.root.Value.Compare(testVal(9)) != 0 {
		t.Error("Failed the left right delete case")
	}
}

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

func TestInsertRoot(t *testing.T) {
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

func TestInsertRightRight(t *testing.T) {
	root := &Node{Height: 1, Value: testVal(10)}
	r1 := &Node{Height: 0, Value: testVal(15)}
	root.right = r1
	tree := &Tree{root: root, Size: 2}
	added := tree.Insert(testVal(20))
	if !added || tree.root.Value.Compare(testVal(15)) != 0 {
		t.Error("Insert right right didn't work")
	}
}

func TestInsertRightLeft(t *testing.T) {
	root := &Node{Height: 1, Value: testVal(10)}
	r1 := &Node{Height: 0, Value: testVal(15)}
	root.right = r1
	tree := &Tree{root: root, Size: 2}
	added := tree.Insert(testVal(13))
	if !added || tree.root.Value.Compare(testVal(13)) != 0 {
		tree.printTree()
		t.Error("Insert right left didn't work")
	}
}

func TestInsertLeftRight(t *testing.T) {
	root := &Node{Height: 1, Value: testVal(10)}
	l1 := &Node{Height: 0, Value: testVal(5)}
	root.left = l1
	tree := &Tree{root: root, Size: 2}
	added := tree.Insert(testVal(7))
	if !added || tree.root.Value.Compare(testVal(7)) != 0 {
		tree.printTree()
		t.Error("Insert left right didn't work")
	}
}

func TestInsertLeftLeft(t *testing.T) {
	root := &Node{Height: 1, Value: testVal(10)}
	l1 := &Node{Height: 0, Value: testVal(5)}
	root.left = l1
	tree := &Tree{root: root, Size: 2}
	added := tree.Insert(testVal(0))
	if !added || tree.root.Value.Compare(testVal(5)) != 0 {
		t.Error("Insert left left didn't work")
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

// MAXNODE

func TestMaxNode(t *testing.T) {
	tree := buildTree()
	val := 17
	if value, ok := tree.MaxNode(); !ok && value.Compare(testVal(val)) != 0 {
		t.Error("Could not find the max node")
	}
}

// MINNODE

func TestMin(t *testing.T) {
	tree := buildTree()
	val := 3
	if value, ok := tree.MinNode(); !ok && value.Compare(testVal(val)) != 0 {
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
	root := &Node{Height: 3, Value: testVal(10)}

	// create the left branch
	l1 := &Node{Height: 1, Value: testVal(7)}
	l2 := &Node{Height: 0, Value: testVal(3)}
	l3 := &Node{Height: 0, Value: testVal(9)}
	l1.left = l2
	l1.right = l3
	root.left = l1

	// create the right branch
	r1 := &Node{Height: 2, Value: testVal(15)}
	r2 := &Node{Height: 0, Value: testVal(12)}
	r3 := &Node{Height: 1, Value: testVal(17)}
	r4 := &Node{Height: 0, Value: testVal(16)}
	r1.left = r2
	r1.right = r3
	r3.left = r4
	root.right = r1

	tree := &Tree{root: root, Size: 7}
	return tree
}

// convenience type
type testVal float64

// the implementation of Comparer
func (t testVal) Compare(test Comparer) int {
	// assert type (ignore any errors for testing)
	otherTest := test.(testVal)
	if t < otherTest {
		return -1
	} else if t > otherTest {
		return 1
	} else {
		return 0
	}
}
