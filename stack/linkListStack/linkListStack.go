package linkListStack

import "fmt"

type Item interface {}

type Node struct {
	Data Item
	Next *Node
}

type Stack struct {
	headNode *Node
}

//push Stack item
func (stack *Stack) Push(item Item) {
	newNode := &Node{Data: item}
	newNode.Next = stack.headNode
	stack.headNode = newNode
}

//pop Item from stack
func (stack *Stack) Pop() Item {
	if stack.headNode == nil {
		fmt.Println("栈已空")
		return nil
	}

	item := stack.headNode.Data
	stack.headNode = stack.headNode.Next
	return item
}

func (stack *Stack) Traverse()  {
	if stack.headNode == nil {
		fmt.Println("栈已空")
		return
	}

	currentNode := stack.headNode
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
}