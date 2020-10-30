package delNNode

import "fmt"

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
func (list List) Traverse() {
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

//获取链表长度
func (list List) Length() int {
	currentNode := list.headNode
	count := 0
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}

	return count
}

//删除单向链表中倒数第N个节点
//方法一：：快慢指针法

func (list *List) DelLastNNode1(lastN int) {
	if lastN > list.Length() || lastN <= 0 {
		fmt.Println("输入的待删结点位置不合法")
		return
	}

	//删除尾结点
	if lastN == 1 {
		list.RemoveLastNode()
		return
	}

	//删除头结点
	if lastN == list.Length() {
		//删除链表头结点
		list.headNode = list.headNode.Next
		return
	}

	lowNode := list.headNode
	fastNode := list.headNode
	prevNode := list.headNode
	fastStep := lastN-1
	for fastNode.Next != nil {
		if fastStep > 0 {
			fastNode = fastNode.Next
			fastStep--
			continue
		}
		fastNode = fastNode.Next
		prevNode = lowNode
		lowNode = lowNode.Next
	}
	prevNode.Next = lowNode.Next
}




//方法二：结点位置和结点数量的关系法
func (list *List) DelLastNNode2(lastN int) {
	if lastN > list.Length() || lastN <= 0{
		fmt.Println("输入的待删结点位置不合法")
		return
	}

	length := list.Length()
	position := length-lastN+1

	if position == 1 {
		//删除链表头结点
		list.headNode = list.headNode.Next
	} else if position == length {
		//删除链表的尾结点
		list.RemoveLastNode()
	} else {
		//删除链表中指定位置的结点
		list.RemovePosition(position)
	}
}

//删除链表尾结点
func (list *List) RemoveLastNode() {
	if list.headNode == nil {
		return
	}

	currentNode := list.headNode
	for currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = nil
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

