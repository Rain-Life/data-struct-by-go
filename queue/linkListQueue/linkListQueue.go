package linkListQueue

import "fmt"

type Item interface {}

type Node struct {
	Data Item
	Next *Node
}

type Queue struct {
	headNode *Node
}

func (queue *Queue) Enqueue(data Item) {
	node := &Node{Data: data}
	if queue.headNode == nil {
		queue.headNode = node
	} else {
		currentNode := queue.headNode
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = node
	}
}

func (queue *Queue) Dequeue() Item {
	if queue.headNode == nil {
		fmt.Println("队列空了")
		return nil
	}

	item := queue.headNode.Data
	queue.headNode = queue.headNode.Next

	return item
}

func (queue *Queue) Traverse() {
	if queue.headNode == nil {
		fmt.Println("队列空的")
		return
	}

	currentNode := queue.headNode
	for currentNode.Next != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Printf("%v\t", currentNode.Data)
}

