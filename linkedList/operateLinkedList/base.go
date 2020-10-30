package operateLinkedList

type Object struct {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

