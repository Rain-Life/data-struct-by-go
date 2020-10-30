package node

import "fmt"

//单向链表
type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

//判断链表是否为空
func (list *List) IsEmpty() bool  {
	if list.headNode == nil {
		return true
	}

	return false
}

//获取链表的长度
func (list *List) Length() int {
	currentNode := list.headNode
	count := 0
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}

	return count
}

//从链表的头部开始添加结点
func (list *List) AddFromHead(data Object) *Node {
	node := &Node{Data:data}
	if list.IsEmpty() {
		list.headNode = node
		return node
	}
	node.Next = list.headNode
	list.headNode = node

	return node
}

//从链表的尾部添加结点
func (list *List) Append(data Object) {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.headNode = node
	} else {
		currentNode := list.headNode
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = node
	}
}

//向单链表中指定位置添加结点
func (list *List) Insert(position int, data Object) {
	if position <= 1 {
		list.AddFromHead(data)
	} else if position > list.Length() {
		list.Append(data)
	} else {
		pre := list.headNode
		count := 1
		for count < (position-1) {
			pre = pre.Next
			count++
		}//循环退出以后pre刚好在position-1的位置
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//删除链表的头结点
func (list *List) RemoveHeadNde() Object {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return nil
	}

	currentNode := list.headNode
	if list.headNode.Next == nil {
		list.headNode = nil
		return currentNode.Data
	}
	list.headNode = currentNode.Next

	return currentNode.Data
}

//删除链表的尾结点
func (list *List) RemoveLastNode() Object {
	if list.IsEmpty() {
		return "链表为空"
	}

	currentNode := list.headNode
	for currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}

	data := currentNode.Next.Data
	currentNode.Next = nil
	return data
}

//删除链表中指定值的节点
func (list *List) Remove(data Object)  {
	pre := list.headNode
	if pre.Data == data{
		list.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

//删除链表中指定位置的元素
func (list *List) RemovePosition(position int)  {
	pre := list.headNode
	if position <= 1 {
		list.headNode = nil
	} else if position > list.Length() {
		fmt.Println("超出链表长度")
	} else {
		count :=1
		for count != position-1 && pre.Next != nil {
			count++
			pre = pre.Next
		}

		pre.Next = pre.Next.Next
	}
}

//查询链表中是否包含某一个结点
func (list *List) Contain(data Object) bool {
	currentNode := list.headNode
	for currentNode != nil {
		if currentNode.Data == data {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

//遍历链表的每一个结点
func (list *List) Traverse()  {
	if list.IsEmpty() {
		fmt.Println("链表为空")
	}
	currentNode := list.headNode
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
}
