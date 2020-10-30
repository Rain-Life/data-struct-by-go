package main

import (
	"fmt"
	"learnGo/linkedList/operateLinkedList/reverseList"
)

func main() {
	list := reverseList.List{}

	fmt.Println("1、就地反转法")
	fmt.Println()

	data := []reverseList.Object{1,2,3,4,5}
	list.CreateList(data)
	fmt.Println("反转前")
	list.Traverse()
	fmt.Println()

	list.ReverseList()
	fmt.Println("反转后")
	list.Traverse()
	fmt.Println()
	fmt.Println()

	fmt.Println("2、头插法反转链表")
	fmt.Println()

	newList := reverseList.List{}
	newData := []reverseList.Object{"a", "b", "c", "d", "e"}
	newList.CreateList(newData)
	fmt.Println("反转前")
	newList.Traverse()
	fmt.Println()

	newList.ReverseListHead()
	fmt.Println()
	fmt.Println()
}
