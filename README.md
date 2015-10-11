# avl-go

[![Build Status](https://travis-ci.org/reedcwilson/avl-go.svg?branch=master)](https://travis-ci.org/reedcwilson/avl-go)

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

### Tree

    type Tree struct {
      Size int
    }

===

#### `func Clear()`

Empties the tree of all its elements.

**Example**

    tree.Clear() // the size should now be zero

===

#### `func Delete(node *Node) bool`

Deletes the element in the tree that matches the given node and re-balances. If
no element matches the node, then Delete returns false.

**Example**

    node := new(avl.Node{ 3 })
    if !tree.Delete(node) {
      fmt.Println("Could not find a node matching your given node")
    }

===

#### `func Find(node *Node) (*Node, bool)`

Returns the element matching the given node or false if no element matches.

**Example**

    node, ok := tree.Find(new(avl.Node{ 3 }))
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

#### `func Insert(node *Node)`

Adds the given node into the tree and performs any necessary balancing.

**Example**

    node := new(avl.Node{ 3 })
    tree.Insert(node)

===

#### `func Max() (*Node, bool)`

Returns the maximum element in the tree. If the tree is empty then it returns
false.

**Example**

    node, err := tree.Max()
    if err {
      fmt.Println("The tree must be empty")
    }

===

#### `func Min() (*Node, bool)`

Returns the minimum element in the tree. If the tree is empty then it returns
false

**Example**

    node, err := tree.Min()
    if err {
      fmt.Println("The tree must be empty")
    }


## Contribute

Fork the repository and clone to the appropriate directory in your GOPATH (i.e.
`$GOPATH/github.com/{your username}/avl-go`)

    cd $GOPATH/github.com/{your username}/avl-go && go install

Ensure that the tests pass with:

    go test
