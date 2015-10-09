package avl

import (
  "errors"
)

type Node struct {
}

type Tree struct {
}

func Clear() {
}

func Delete(node Node) bool {
  return false
}

func Find(node Node) (Node, error) {
  return Node{}, errors.New("Not Implemented")
}

func Height() int {
  return 0
}

func Insert(node Node)  {
}

func Largest() (Node, error) {
  return Node{}, errors.New("Not Implemented")
}

func Size() int {
  return 0
}

func Smallest() (Node, error) {
  return Node{}, errors.New("Not Implemented")
}
