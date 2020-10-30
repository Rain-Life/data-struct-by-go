package main

import (
	"fmt"
	"learnGo/linkedList/loopLinkList"
)

func print(list loopLinkList.List)  {
	fmt.Printf("链表长度：%d\n", list.Length())
	fmt.Println("遍历链表所有结点：")
	list.Traverse()
	fmt.Println()
	fmt.Println()
}

func main() {
	list := loopLinkList.List{}

	fmt.Println("++++++++1、判断链表是否为空++++++++")
	fmt.Printf("链表是否为空：%v\n", list.IsEmpty())
	print(list)

	fmt.Println("++++++++2、获取链表长度++++++++")
	fmt.Printf("获取链表长度：%d\n", list.Length())
	print(list)

	fmt.Println("++++++++3、向循环链头部添加结点++++++++")
	list.AddFromHead("firstNode")
	print(list)

	fmt.Println("++++++++4、向循环链尾部添加结点++++++++")
	list.Append("lastNode")
	print(list)

	fmt.Println("++++++++5、向循环链指定位置添加结点++++++++")
	list.Insert(1,"secondNode")
	print(list)

	fmt.Println("++++++++6、删除循环链的头结点++++++++")
	list.RemoveHeadNde()
	print(list)

	fmt.Println("++++++++7、删除循环链的尾结点++++++++")
	list.RemoveLastNode()
	print(list)

	fmt.Println("++++++++8、查找循环链表中指定结点++++++++")
	fmt.Printf("是否包含secondNode节点：%v\n", list.Contain("secondNode"))
	print(list)

	fmt.Println("++++++++9、删除循环链表中指定位置的结点++++++++")
	list.RemovePosition(1)
	print(list)

	fmt.Println("++++++++10、删除循环链表中指定值的结点++++++++")
	list.Remove("lastNode")
	print(list)

}
