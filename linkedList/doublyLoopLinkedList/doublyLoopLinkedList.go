package doublyLoopLinkedList

import (
	"fmt"
)

type Object interface {}

type Node struct {
	Data Object
	Prev,Next *Node
}

type List struct {
	headNode *Node
}

//判断双向循环链表是否为空
func (list *List) IsEmpty() bool {
	if list.headNode == nil {
		return true
	}

	return false
}

//遍历双向循环链表
func (list *List) Traverse() {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	if list.headNode.Next == nil {//兼容双向循环链表只有一个结点的情况
		fmt.Printf("%v\t", list.headNode.Data)
		return
	}

	currentNode := list.headNode
	for currentNode.Next != list.headNode {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Printf("%v\t", currentNode.Data)
}

//获取双向循环链表的长度
func (list *List) Length() int {
	if list.IsEmpty() {
		return 0
	}

	count := 1
	currentNode := list.headNode
	for currentNode.Next != list.headNode {
		count++
		currentNode = currentNode.Next
	}

	return count
}

//从双向循环链表头部添加结点(一定一定要画图，比较好理解)
func (list *List) AddFromHead(data Object) *Node {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.headNode = node
		return node
	}

	currentNode := list.headNode

	if currentNode.Next == nil { //考虑只有一个节点的时候
		node.Prev = currentNode
		currentNode.Next = node
		node.Next = currentNode
		currentNode.Prev = node
		list.headNode = node
		return node
	}

	node.Prev = currentNode.Prev
	currentNode.Prev.Next = node
	currentNode.Prev = node
	node.Next = currentNode

	list.headNode = node

	return node
}

//从双向循环链表尾部添加结点
func (list *List) Append(data Object) *Node {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.headNode = node
		return node
	}

	currentNode := list.headNode
	currentNode.Prev.Next = node
	node.Prev = currentNode.Prev
	currentNode.Prev = node
	node.Next = currentNode

	return node
}

//向双向循环链表的指定位置插入结点
func (list *List) Insert(position int, data Object) {
	if position <=1 {
		list.AddFromHead(data)
	} else if position > list.Length() {
		list.Append(data)
	} else {
		node := &Node{Data: data}
		currentNode := list.headNode
		count := 1
		for count != position {
			count++
			currentNode = currentNode.Next
		}

		//将待插入的节点插入到当前节点的前边即可(这块的逻辑其实和双向链表的在中间位子插入结点逻辑一样)
		currentNode.Prev.Next = node
		node.Prev = currentNode.Prev
		currentNode.Prev = node
		node.Next = currentNode
	}
}

//删除双向循环链表头结点
func (list *List) RemoveHeadNde() Object {
	if list.IsEmpty() {
		fmt.Println("双向循环链表为空")
		return ""
	}

	currentNode := list.headNode
	if currentNode.Next == nil {//只有一个结点的情况
		list.headNode = nil
		return currentNode.Data
	}

	currentNode.Prev.Next = currentNode.Next
	currentNode.Next.Prev = currentNode.Prev
	list.headNode = currentNode.Next

	return currentNode.Data
}

//删除双向循环链表尾结点
func (list *List) RemoveLastNode() Object {
	if list.IsEmpty() {
		fmt.Println("双向循环链表为空")
		return ""
	}

	currentNode := list.headNode
	if currentNode.Next == nil {//只有一个结点的情况
		list.headNode = nil
		return list.headNode.Data
	}

	lastNode := list.headNode.Prev
	lastNode.Prev.Next = currentNode
	currentNode.Prev = lastNode.Prev

	return lastNode.Data
}

//删除双向循环链表中指定值的结点
func (list *List) Remove(data Object) bool {
	if list.IsEmpty() {
		fmt.Println("双向循环链表为空")
		return false
	}

	//如果链表只有一个节点
	if list.headNode.Next == nil {
		if list.headNode.Data == data {
			list.headNode = nil
			return true
		}
		return false
	}

	currentNode := list.headNode
	//如果值刚好等于第一个节点的值
	if currentNode.Data == data {
		list.RemoveHeadNde()
		return true
	}

	//如果值刚好等于最后一个结点的值
	for currentNode.Next != list.headNode {
		if currentNode.Data == data {
			//删除中间节点
			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
			return true
		} else {
			currentNode = currentNode.Next
		}
	}

	if currentNode.Data == data { //刚好是最后一个结点的时候
		list.RemoveLastNode()
		return true
	}

	return false
}

//删除双向循环链表中指定位置的结点
func (list *List) RemovePosition(position int) {

	if list.IsEmpty() {
		fmt.Println("双向循环链表为空")
		return
	}

	if position <= 1 {
		list.RemoveHeadNde()
	} else if position > list.Length() {
		list.RemoveLastNode()
	} else {
		currentNode := list.headNode
		count := 1
		for count != position {
			count++
			currentNode = currentNode.Next
		}

		currentNode.Prev.Next = currentNode.Next
		currentNode.Next.Prev = currentNode.Prev
	}
}

//查询双向循环链表中是否包含某一个结点(跟循环链表的查找逻辑一样)
func (list *List) Contain(data Object) bool {
	if list.IsEmpty() {
		fmt.Println("双向循环链表为空")
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


