package avl

type Stack []*Node

func (q *Stack) Push(n *Node) {
  *q = append(*q, n)
}

func (q *Stack) Pop() (n *Node) {
  x := q.Len() - 1
  n = (*q)[x]
  *q = (*q)[:x]
  return
}
func (q *Stack) Len() int {
  return len(*q)
}
