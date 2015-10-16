package avl

// Every value that is contained in a node must implement Comparer where < 0
// indicates less than, 0 indicates equal and > 1 indicates greater than
type Comparer interface {
  Compare(Comparer) int
}

