package main

import (
	"base01/struct/queue"
	"fmt"
)

func main() {
	//var root tree.Node
	//
	//root = tree.Node{Value: 3}
	//
	//root.Left = &tree.Node{}
	//
	//root.Right = &tree.Node{10, nil, nil}
	//
	//root.Right.Left = new(tree.Node)
	//
	//root.Left.Right = tree.CreateNode(6)
	//
	//root.Right.Left.SetValue(8)
	//
	//root.Traverse()

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	q.Pop()
	q.Pop()
	q.Pop()
	if q.IsEmpty() {
		fmt.Println("queue is none")
	} else {
		fmt.Println("queue is ", q)
	}

}
