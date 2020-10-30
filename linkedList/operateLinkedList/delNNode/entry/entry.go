package main

import (
	"fmt"
	"learnGo/linkedList/operateLinkedList/delNNode"
)

func main() {
	list := delNNode.List{}

	fmt.Println("1、快慢指针法")
	fmt.Println()

	data := []delNNode.Object{"e","d","c","b","a"}
	list.CreateList(data)
	fmt.Println("删除倒数第N个结点前")
	list.Traverse()
	fmt.Println()

	list.DelLastNNode1(5)
	fmt.Println("删除倒数第N个结点后")
	list.Traverse()

	//fmt.Println("2、结点位置和结点数量的关系法")
	//fmt.Println()
	//
	//data := []delNNode.Object{1,2,3,4,5}
	//list.CreateList(data)
	//fmt.Println("删除倒数第N个结点前")
	//list.Traverse()
	//fmt.Println()
	//
	//list.DelLastNNode2(4)
	//fmt.Println("删除倒数第N个结点后")
	//list.Traverse()

}
