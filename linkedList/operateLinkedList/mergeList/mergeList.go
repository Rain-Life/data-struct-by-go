package mergeList

import "fmt"

type Object interface {}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	HeadNode *Node
}

//生成单链表
func (list *List) CreateList(data []Object) {
	count := len(data)

	if count < 0 || count > 1000 {
		fmt.Println("参数不合法")
		return
	} else if count == 1 {
		node := &Node{Data: data[0]}
		list.HeadNode = node
	} else {
		for _, value := range data {
			node := &Node{Data: value}
			//为了方便，这里使用头插法
			node.Next = list.HeadNode
			list.HeadNode = node
		}
	}
}

//遍历链表
func (list List) Traverse() {
	if list.HeadNode == nil {
		fmt.Println("链表为空")
		return
	}

	currentNode := list.HeadNode
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
	}
}

//方法一：常规方法合并链表
//解题思路：https://juejin.im/post/6844903855965077518
func (list *List) MergeLinkedList(list1 List, list2 List) {
	if list1.HeadNode == nil && list2.HeadNode == nil {
		fmt.Println("两个链表均为空")
		return
	} else if list1.HeadNode == nil {
		list.HeadNode = list2.HeadNode
		return
	} else if list2.HeadNode == nil {
		list.HeadNode = list1.HeadNode
		return
	}

	curr1 := list1.HeadNode
	curr2 := list2.HeadNode
	if curr1.Data.(int) < curr2.Data.(int) {
		list.HeadNode = curr1
		curr1 = curr1.Next
	} else {
		list.HeadNode = curr2
		curr2 = curr2.Next
	}
	newHead := list.HeadNode
	currNew := newHead

	for curr1 != nil && curr2 != nil {
		if curr1.Data.(int) < curr2.Data.(int) {
			currNew.Next = curr1
			curr1 = curr1.Next
		} else {
			currNew.Next = curr2
			curr2 = curr2.Next
		}
		currNew = currNew.Next
	}

	if curr1 == nil {
		currNew.Next = curr2
	}
	if curr2 == nil {
		currNew.Next = curr1
	}
}

//方法二：通过递归的方法合并链表
//解题思路：https://juejin.im/post/6844903855965077518
func RecursionMergeList(head1 *Node, head2 *Node) *Node {
	newNode := &Node{}
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	} else {
		if head1.Data.(int) < head2.Data.(int) {
			newNode = head1
			newNode.Next = RecursionMergeList(head1.Next, head2)
		} else {
			newNode = head2
			newNode.Next = RecursionMergeList(head1, head2.Next)
		}
		return newNode
	}
}






