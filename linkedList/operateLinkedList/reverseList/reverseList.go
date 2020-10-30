package reverseList

import "fmt"

//单链表反转

type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//生成单链表
func (list *List) CreateList(data []Object) {
	count := len(data)

	if count < 0 || count > 1000 {
		fmt.Println("参数不合法")
		return
	} else if count == 1 {
		node := &Node{Data: data[0]}
		list.headNode = node
	} else {
		for _, value := range data {
			node := &Node{Data: value}
			//为了方便，这里使用头插法
			node.Next = list.headNode
			list.headNode = node
		}
	}
}


//遍历链表
func (list *List) Traverse() {
	if list.headNode == nil {
		fmt.Println("链表为空")
		return
	}

	currentNode := list.headNode
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
}

//参考：https://www.cnblogs.com/byrhuangqiang/p/4311336.html
//单链表反转：就地反转法
func (list *List) ReverseList() {
	if list.headNode == nil {
		fmt.Println("链表为空")
		return
	}

	newHead := &Node{}
	newHead.Next = list.headNode
	prevNode := newHead.Next
	currentNode := prevNode.Next

	for currentNode != nil {
		prevNode.Next = currentNode.Next
		currentNode.Next = newHead.Next
		newHead.Next = currentNode
		currentNode = prevNode.Next
	}
	list.headNode = newHead.Next
}

//单链表反转：头插法反转链表
func (list *List) ReverseListHead() {
	if list.headNode == nil {
		fmt.Println("链表为空")
		return
	}

	newList := &List{}
	currentNode := list.headNode
	nextNode := currentNode.Next
	for currentNode!=nil {
		if newList.headNode == nil {
			newList.headNode = currentNode
			newList.headNode.Next = nil
			currentNode = nextNode
			continue
		}
		nextNode = currentNode.Next
		currentNode.Next = newList.headNode
		newList.headNode = currentNode
		currentNode = nextNode
	}

	fmt.Println("反转后")
	newList.Traverse()
}
