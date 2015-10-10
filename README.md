# avl-go

[![Build Status](https://travis-ci.org/reedcwilson/alv-go.svg?branch=master)](https://travis-ci.org/reedcwilson/alv-go)

This is an exercise to learn golang. 

## Usage

Place an import for avl-go in your file:

    import (
      "fmt"
      "github.com/reedcwilson/avl-go"
    )

And (assuming you are in the directory where your file is) use go to fetch it
for you.

    go get

*NOTE*: Although the repository is `avl-go` the package name is `avl` 

You can then use the tree like so:

    tree := new(avl.Tree)
    tree.Insert(new(avl.Node{ 3 }))


## API

### Node

#### `func Compare(node Node)`

Compares itself with the given node. Compare returns zero if the elements match,
a negative number if the given node is less than the current node and a positive
number if greater.

**Example**

    n1 := new(avl.Node{ 1 })
    n2 := new(avl.Node{ 2 })
    result := n1.Compare(n2)  // -1

### Tree

#### `func Clear()`

Empties the tree of all its elements.

**Example**

    tree.Clear() // the size should now be zero

===

#### `func Delete(node Node) bool`

Deletes the element in the tree that matches the given node and re-balances. If
no element matches the node, then Delete returns false.

**Example**

    node := new(tree.Node{ 3 })
    if !tree.Delete(node) {
      fmt.Println("Could not find a node matching your given node")
    }

===

#### `func Find(node Node) (Node, error)`

Returns the element matching the given node. If no element matches then an error
is returned with a zero Node.

**Example**

    node, err := tree.Find(new(tree.Node{ 3 }))
    if (err) {
      fmt.Println("Could not find a node matching your given node")
    }

===

#### `func Height() int`

Returns the number of levels in the tree.

**Example**

    height := tree.Height()
    if height > 100 {
      fmt.Println("That is a big tree at %d layers", height)
    }

===

#### `func Insert(node Node)`

Adds the given node into the tree and performs any necessary balancing.

**Example**

    node := new(tree.Node{ 3 })
    tree.Insert(node)

===

#### `func Largest() (Node, error)`

Returns the largest element in the tree. If the tree is empty then it returns an
error along with a zero Node.

**Example**

    node, err := tree.Largest()
    if err {
      fmt.Println("The tree must be empty")
    }

===

#### `func Size() int`

Returns the total number of elements in the tree.

**Example**

    size := tree.Size()
    if size < 5 {
      fmt.Println("One day you might be more than a sapling")
    }

===

#### `func Smallest() (Node, error)`

Returns the smallest element in the tree. If the tree is empty then it returns
an error along with a zero Node.

**Example**

    node, err := tree.Smallest()
    if err {
      fmt.Println("The tree must be empty")
    }


## Contribute

Fork the repository and clone to the appropriate directory in your GOPATH (i.e.
`$GOPATH/github.com/{your username}/avl-go`)

    cd $GOPATH/github.com/{your username}/avl-go && go install

Ensure that the tests pass with:

    go test
