package main

import (
	"fmt"
	"learnGo/linkedList/operateLinkedList/josephCircle"
)

func main() {
	list := josephCircle.List{}
	data := []josephCircle.Object{1,2,3,4,5}
	fmt.Println("创建单向循环链表")
	list.CreateLoopList(data)
	list.Traverse()
	fmt.Println()

	fmt.Println("单向循环链表解决约瑟夫问题")
	arr := list.DealJosephCircle(3)
	for _, value := range arr {
		fmt.Printf("%v\t", value)
	}
	fmt.Println()
}
