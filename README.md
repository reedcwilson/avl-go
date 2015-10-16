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

You will have to make sure that the elements you wish to place in the tree
implement a certain interface.

    type Comparer interface {
      Compare(Comparer) float64
    }

You should return a negative number if the current object is less than the
argument, positive for greater than, and 0 for equal.

For raw types you can bind a method to a declared type.

    type myFloat float64
    func (f1 myFloat) Compare(value Comparer) {
      f2, ok := value.(myFloat) // type assertion
      ...
    }

You can then use the tree like so:

    tree := new(avl.Tree)
    tree.Insert(myFloat(24))


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

#### `func Delete(value Comparer) bool`

Deletes the element in the tree that matches the given element and re-balances. If
no element matches the value, then Delete returns false.

**Example**

    if !tree.Delete(myFloat(6)) {
      fmt.Println("Could not find an element matching your given value")
    }

===

#### `func Find(value Comparer) (Comparer, bool)`

Returns the element matching the given value or false if no element matches.

**Example**

    value, ok := tree.Find(myFloat(4))
    if !ok {
      fmt.Println("Could not find an elemtn matching your given value")
    }

===

#### `func Height() int`

Returns the depth of the tree.

**Example**

    height := tree.Height()
    if height > 100 {
      fmt.Println("That is a big tree at %d layers", height)
    }

===

#### `func Insert(value Comparer)`

Adds the given element into the tree and performs any necessary balancing.

**Example**

    tree.Insert(myFloat(1))

===

#### `func Max() (Comparer, bool)`

Returns the maximum element in the tree. If the tree is empty then it returns
false.

**Example**

    value, ok := tree.Max()
    if !ok {
      fmt.Println("The tree appears to be empty")
    }

===

#### `func Min() (Comparer, bool)`

Returns the minimum element in the tree. If the tree is empty then it returns
false

**Example**

    value, ok := tree.Min()
    if !ok {
      fmt.Println("The tree appears to be empty")
    }


## Contribute

Fork the repository and clone to the appropriate directory in your GOPATH (i.e.
`$GOPATH/github.com/{your username}/avl-go`)

    cd $GOPATH/github.com/{your username}/avl-go && go install

Ensure that the tests pass with:

    go test
