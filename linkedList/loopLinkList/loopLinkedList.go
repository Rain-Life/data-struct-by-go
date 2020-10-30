package loopLinkList

import "fmt"

//循环链表
type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//判断循环链表是否为空（与单链表的实现没有区别）
func (list *List) IsEmpty() bool {
	if list.headNode == nil {
		return true
	}

	return false
}

//获取循环链表的长度（与单链表的获取长度区别在于循环的终止条件）
func (list *List) Length() int {
	if list.headNode == nil {
		return 0
	}

	currentNode := list.headNode
	count := 1
	for currentNode.Next != list.headNode {
		count++
		currentNode = currentNode.Next
	}

	return count
}

//向循环链头部添加结点
func (list *List) AddFromHead(data Object) {
	node := &Node{Data: data}
	//链表为空的情况
	if list.IsEmpty() {
		list.headNode = node
		list.headNode.Next = list.headNode //单链表没有这一步
	} else {
		currentNode := list.headNode
		for currentNode.Next != list.headNode { //需要先找到最后一个结点
			currentNode = currentNode.Next
		}

		node.Next = list.headNode
		currentNode.Next = node //单链表没有这一步操作，将最后一个节点的next指向头结点
		list.headNode = node
	}
}


//向循环链表的尾部添加结点
func (list *List) Append(data Object) {
	node := &Node{Data: data}

	if list.IsEmpty() {
		list.headNode = node
		list.headNode.Next = node
	} else {
		currentNode := list.headNode
		for currentNode.Next != list.headNode {
			currentNode = currentNode.Next
		}

		currentNode.Next = node
		node.Next = list.headNode //单链表没有这一步操作（让最后一个节点指向头结点）
	}
}

//向循环链表的指定位置添加结点（跟单链表是一样的）
func (list *List) Insert(position int, data Object) {
	if position <= 1 {
		list.AddFromHead(data)
	} else if position > list.Length() {
		list.Append(data)
	} else {
		currentNode := list.headNode
		count := 1
		for count < position-1 {
			currentNode = currentNode.Next
			count++
		}

		node := &Node{Data: data}
		node.Next = currentNode.Next
		currentNode.Next = node
	}
}

//删除循环链表的头结点
func (list *List) RemoveHeadNde() {
	if list.IsEmpty() {
		fmt.Println("链表是空的")
		return
	}

	currentNode := list.headNode
	lastNode := list.headNode
	//考虑循环链表只有一个结点的情况
	if currentNode.Next == list.headNode {
		list.headNode = nil
		return
	}

	for lastNode.Next != list.headNode {
		lastNode = lastNode.Next
	}

	list.headNode = currentNode.Next
	lastNode.Next = list.headNode
}

//删除循环链表的尾结点
func (list *List) RemoveLastNode() {
	if list.IsEmpty() {
		fmt.Println("链表是空的")
		return
	}

	currentNode := list.headNode
	//考虑循环链表只有一个结点的情况
	if currentNode.Next == list.headNode {
		list.headNode = nil
		return
	}

	for currentNode.Next.Next != list.headNode {
		currentNode = currentNode.Next
	}

	currentNode.Next = list.headNode
}

//删除循环链表中指定位置的节点 (需考虑删除的结点是不是第一个或最后一个)
func (list *List) RemovePosition(position int) {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	if position <= 1 {
		list.RemoveHeadNde()
	} else if position > list.Length() {
		list.RemoveLastNode()
	} else {
		currentNode := list.headNode
		count := 1
		if count != position-1 && currentNode.Next != list.headNode{
			count++
			currentNode = currentNode.Next
		}

		currentNode.Next = currentNode.Next.Next
	}
}

//删除循环链表中指定值的结点
func (list *List) Remove(data Object) {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	currentNode := list.headNode

	//删除的结点为头结点时
	if currentNode.Data == data {
		list.RemoveHeadNde()
		return
	}

	for currentNode.Next != list.headNode {
		if currentNode.Next.Data == data {
			currentNode.Next = currentNode.Next.Next
			return
		} else {
			currentNode = currentNode.Next
		}
	}

	if currentNode.Next == list.headNode {
		list.RemoveLastNode()
	}
}

//查找循环链表中指定结点
func (list *List) Contain(data Object) bool {
	if list.IsEmpty() {
		return false
	}

	currentNode := list.headNode
	for currentNode.Next != list.headNode {
		if currentNode.Data == data {
			return true
		}

		currentNode = currentNode.Next
	}

	if currentNode.Data == data {
		return true
	}

	return false
}

//遍历循环链表
func (list *List) Traverse() {
	if list.IsEmpty() {
		fmt.Println("循环链表为空")
		return
	}

	currentNode := list.headNode
	if currentNode.Next == list.headNode { //兼容循环链表只有一个结点的情况
		fmt.Printf("%v\t", currentNode.Data)
		return
	}

	for currentNode.Next != list.headNode {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Printf("%v\t", currentNode.Data)
}


