package main

import (
	"fmt"
	"learnGo/linkedList/operateLinkedList/checkRing"
)

func main() {
	list := checkRing.List{}

	fmt.Println("1、生成无环单链表")
	data := []checkRing.Object{"e","d","c","b","a"}
	list.CreateList(data)
	list.Traverse()
	fmt.Println()
	fmt.Println("链表环的检测结果：", list.CheckRing())
	fmt.Println()

	ringList := checkRing.List{}
	fmt.Println("2、生成有环的单链表")
	dataRing := []checkRing.Object{"h","g","f","e","d","c","b","a"}
	ringList.CreateRingList(dataRing)
	ringList.Traverse()
	fmt.Println()
	fmt.Println("链表环的检测结果：", ringList.CheckRing())
	fmt.Println()
}
