package checkRing

import (
	"fmt"
)

//单向链表中环的检测

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

func (list *List) CreateRingList(data []Object) {
	count := len(data)
	if count <= 1 {
		fmt.Println("该链表的结点个数无法形成环")
		return
	} else {
		nodeCount := 0
		lastNode := &Node{}
		thirdNode := &Node{}
		for _, value := range data {
			node := &Node{Data: value}
			nodeCount++
			if nodeCount == 1 {
				lastNode = node
			}
			if nodeCount == 3 {
				thirdNode = node
			}
			node.Next = list.headNode
			list.headNode = node
			if nodeCount == count {
				lastNode.Next = thirdNode
			}
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
	count := 0
	for currentNode != nil {
		fmt.Printf("%v\t", currentNode.Data)
		currentNode = currentNode.Next
		count++
		if count > 10 {
			break
		}
	}
}


//快慢指针法

func (list List) CheckRing() bool {
	if list.headNode == nil {
		fmt.Println("链表为空")
		return false
	}

	low := list.headNode
	fast := list.headNode
	for fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next

		//为了防止for中fast.Next出现nil取Next，这里先做个判断
		if fast == nil {
			return false
		}
		if low == fast {
			return true
		}
	}

	return false
}





