package main

import (
	"fmt"
	"learnGo/linkedList/operateLinkedList/mergeList"
)

func main() {
	list1 := &mergeList.List{}
	list2 := &mergeList.List{}
	newList := mergeList.List{}
	data1 := []mergeList.Object{4,2,1}
	data2 := []mergeList.Object{4,3,1}
	list1.CreateList(data1)
	fmt.Println("链表1")
	list1.Traverse()
	fmt.Println()

	list2.CreateList(data2)
	fmt.Println("链表2")
	list2.Traverse()
	fmt.Println()

	//fmt.Println("list1为空，链表合并结果")
	//newList.MergeLinkedList(*list1, *list2)
	//newList.Traverse()
	//fmt.Println()

	//fmt.Println("list2为空，链表合并结果")
	//newList.MergeLinkedList(*list1, *list2)
	//newList.Traverse()

	fmt.Println("两个链表均不为空，链表合并结果")
	newList.MergeLinkedList(*list1, *list2)
	newList.Traverse()

}
