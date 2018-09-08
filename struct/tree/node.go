package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//创建
func CreateNode(value int) *Node {
	//返回局部变量地址
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Print(node.Value)
}

//加*引用传递 不加值传递
func (node *Node) SetValue(value int) {
	node.Value = value
}

//遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Printf("%d ", node.Value)
	node.Right.Traverse()
}
