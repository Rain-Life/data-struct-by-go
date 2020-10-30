package doublyLinkedList

import (
	"fmt"
)

//双向链表
type Object interface {}

type Node struct {
	Data Object
	Prev,Next *Node
}

type List struct {
	headNode *Node
}

//说明
/**
1、如果结点的Next指向为null，则说明是最后一个结点
2、第一个结点的prev指向为空
3、双向链表和单向链表差不多，区别就是删除和添加结点的时候更方便了，因为很方便的可以获取到前一个结点
*/

//判断双向链表是否为空
func (list *List) IsEmpty() bool {
	if list.headNode == nil {
		return true
	}

	return false
}

//遍历双向链表（跟遍历单链表一模一样）
func (list *List) Traverse() {
	if list.IsEmpty() {
		fmt.Println("双线链表为空")
		return
	}

	currentNode := list.headNode
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
}

//获取双向链表的长度（跟获取单链表长度一模一样）
func (list *List) Length() int {
	if list.IsEmpty() {
		return 0
	}

	count := 0
	currentNode := list.headNode
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}

	return count
}

//从双向链表头部开始增加结点
func (list *List) AddFromHead(data Object) *Node {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.headNode = node
		return node
	}

	list.headNode.Prev = node
	node.Next = list.headNode
	list.headNode = node

	return node
}

//从双向链表尾部添加结点
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
		node.Prev = currentNode
	}
}

/**
 * 向双向链表的指定位置插入结点
 *
 * 说明：在单向链表中，我是通过找到我要插入的这个结点的前一个结点，然后将要插入的结点，
 * 插入到这个结点的后边（因为插入结点需要找到当前结点的前一个结点，为了避免再次遍历找前一个结点，所以采用了这种方式）
 * 而双向链表就不需要这么做了，找到指定位置的结点，新的插入到它前边就可以了
 */
func (list *List) Insert(position int, data Object) {
	if position <= 1 {
		list.AddFromHead(data)
	} else if position > list.Length() {
		list.Append(data)
	} else {
		currentNode := list.headNode
		count := 1
		for count != position {
			currentNode = currentNode.Next
			count++
		}
		//找到了指定位置的结点，然后将要插入的结点，插到这个节点前边即可(注意顺序，画图最容易理解)
		node := &Node{Data: data}
		node.Next = currentNode
		node.Prev = currentNode.Prev
		currentNode.Prev.Next = node
		currentNode.Prev = node
	}
}

//删除链表头结点
func (list *List) RemoveHeadNde() Object {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return nil
	}

	currentNode := list.headNode
	if currentNode.Next == nil && currentNode.Prev == nil {
		list.headNode = nil
		return currentNode.Data
	}

	list.headNode = currentNode.Next
	currentNode.Prev = nil

	return currentNode.Data
}

//删除链表尾结点
func (list *List) RemoveLastNode() Object {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return nil
	}

	if list.headNode.Next == nil {
		list.headNode = nil
	}
	
	currentNode := list.headNode
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Prev.Next = nil

	return currentNode.Prev.Data
}

//删除双向表中指定值的结点
func (list *List) Remove(data Object)  {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	currentNode := list.headNode
	if list.headNode.Next == nil && list.headNode.Data == data {
		list.headNode = nil
	}

	fmt.Println(data, currentNode.Data, currentNode.Data == data)
	for currentNode != nil {
		if currentNode.Data == data && currentNode == list.headNode {
			list.headNode = currentNode.Next
		} else if currentNode.Data == data && currentNode.Prev != nil {
			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
		} else {
			currentNode = currentNode.Next
		}
	}
}

//删除双向链表中指定位置的结点
func (list *List) RemovePosition(position int) {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	if position <=1 {
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

//查询双向链表中是否包含某一个结点(和单向链表一样)
func (list *List) Contain(data Object) bool {
	if list.IsEmpty() {
		fmt.Println("链表为空")
		return false
	}

	currentNode := list.headNode
	for currentNode != nil {
		if currentNode.Data == data {
			return true
		}

		currentNode = currentNode.Next
	}

	return false
}
