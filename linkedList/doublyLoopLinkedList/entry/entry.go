package main

import (
	"fmt"
	"learnGo/linkedList/doublyLoopLinkedList"
)

func print(list doublyLoopLinkedList.List)  {
	fmt.Printf("链表长度%d\n", list.Length())
	fmt.Println("遍历链表所有结点：")
	list.Traverse()
	fmt.Println()
	fmt.Println()
}

func main() {
	list := doublyLoopLinkedList.List{}

	fmt.Println("++++++++1、判断链表是否为空++++++++")
	fmt.Printf("链表是否为空：%v", list.IsEmpty())
	fmt.Println()
	fmt.Println()

	fmt.Println("++++++++2、向双向循环链表头部添加结点++++++++")
	fmt.Println()
	list.AddFromHead(1)
	list.AddFromHead(2)
	list.AddFromHead(3)
	print(list)

	fmt.Println("++++++++3、向双向循环链表尾部添加结点++++++++")
	fmt.Println()
	list.Append("lastNode")
	print(list)

	fmt.Println("++++++++4、向双向循环链表指定位置添加结点++++++++")
	fmt.Println()
	list.Insert(1,"firstNode")
	print(list)

	fmt.Println("++++++++5、删除双向循环链表头结点++++++++")
	fmt.Println()
	list.RemoveHeadNde()
	print(list)

	fmt.Println("++++++++6、删除双向循环链表尾结点++++++++")
	fmt.Println()
	list.RemoveLastNode()
	print(list)

	fmt.Println("++++++++7、删除双向循环链表中指定值的结点++++++++")
	fmt.Println()
	list.Remove(1)
	print(list)

	fmt.Println("++++++++8、删除双向循环链表中指定位置的结点++++++++")
	fmt.Println()
	list.RemovePosition(list.Length()-1)
	print(list)

	fmt.Println("++++++++9、查询双向循环链表中是否包含某一个结点++++++++")
	fmt.Println()
	fmt.Println("是否包含值为firstNode的结点：", list.Contain("firstNode"))
	print(list)
}
