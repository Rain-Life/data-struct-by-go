package main

import (
	"fmt"
	"learnGo/linkedList/node"
)

func print(list node.List)  {
	fmt.Printf("链表长度%d\n", list.Length())
	fmt.Println("遍历链表所有结点：")
	list.Traverse()
	fmt.Println()
	fmt.Println()
}

func main() {
	list := node.List{}

	//向链表的尾部追加元素
	fmt.Println("++++++++1、向链表尾部追加结点++++++++")
	list.Append(3)
	list.Append(8)
	list.Append(1)
	list.Append(9)

	list.Append("PHP")
	list.Append("Golang")
	list.Append("Java")
	list.Append("JavaScript")
	print(list)

	fmt.Println("++++++++2、判断链表是否为空++++++++")
	fmt.Printf("链表是否为空：%v", list.IsEmpty())
	fmt.Println()
	fmt.Println()

	//向链表的头部添加元素
	fmt.Println("++++++++3、向链表的头部添加一个结点++++++++")
	fmt.Println()
	list.AddFromHead("firstNode")
	print(list)

	fmt.Println("++++++++4、向链表尾部添加结点++++++++")
	fmt.Println()
	list.Append("lastNode")
	print(list)

	fmt.Println("++++++++5、向链指定位置添加结点++++++++")
	fmt.Println()
	list.Insert(11, "thirdNode")
	print(list)

	fmt.Println("++++++++6、删除链表头结点++++++++")
	fmt.Println()
	list.RemoveHeadNde()
	print(list)

	fmt.Println("++++++++7、删除链表尾结点++++++++")
	fmt.Println()
	list.RemoveLastNode()
	print(list)

	fmt.Println("++++++++8、删除链表中指定值的结点++++++++")
	fmt.Println()
	list.Remove(9)
	print(list)

	fmt.Println("++++++++9、删除链表中指定位置的结点++++++++")
	fmt.Println()
	list.RemovePosition(3)
	print(list)

	fmt.Println("++++++++10、查询链表中是否包含某一个结点++++++++")
	fmt.Println()
	res := list.Contain("Golang")
	fmt.Printf("是否存在Golang结点：%v\n", res)
	print(list)
}
