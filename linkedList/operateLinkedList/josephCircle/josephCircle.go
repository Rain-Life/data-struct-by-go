package josephCircle

import "fmt"

type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//遍历链表
func (list List)Traverse()  {
	if list.headNode == nil {
		fmt.Printf("链表为空")
		return
	}

	currentNode := list.headNode
	for currentNode.Next != list.headNode {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Printf("%v\t", currentNode.Data)
}

//生成循环链表
func (list *List)CreateLoopList(data []Object)  {
	if len(data) < 0 {
		fmt.Println("参数不合法")
		return
	}

	currentNode := list.headNode
	for i, value := range data {
		node := &Node{Data: value}
		if i == 0 {
			list.headNode = node
			currentNode = node
			continue
		}
		currentNode.Next = node
		currentNode = node
	}
	currentNode.Next = list.headNode
}

//单向链表解决约瑟夫问题
func (list *List) DealJosephCircle(number int) []interface{} {
	var data []interface{}
	index := 1
	currentNode := list.headNode
	preNode := list.headNode
	for preNode.Next != list.headNode {
		preNode = preNode.Next //刚开始，使preNode指向最后一个节点
	}

	for currentNode.Next != currentNode {
		if index == number {
			//删除结点，其实不用考虑是不是头结点或尾结点
			data = append(data, currentNode.Data)
			preNode.Next = currentNode.Next
			currentNode = preNode.Next
			index = 1

			continue
		} else {
			index++
		}
		preNode = currentNode
		currentNode = currentNode.Next
	}
	data = append(data, currentNode.Data)
	currentNode = nil

	return data
}
