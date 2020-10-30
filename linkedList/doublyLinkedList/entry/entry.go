package main

import (
	"fmt"
	"learnGo/linkedList/doublyLinkedList"
)

func print(list doublyLinkedList.List)  {
	fmt.Printf("链表长度%d\n", list.Length())
	fmt.Println("遍历链表所有结点：")
	list.Traverse()
	fmt.Println()
	fmt.Println()
}

func main() {
	list := doublyLinkedList.List{}

	fmt.Println("++++++++1、判断链表是否为空++++++++")
	fmt.Printf("链表是否为空：%v", list.IsEmpty())
	fmt.Println()
	fmt.Println()

	fmt.Println("++++++++2、向双向链表头部添加元素++++++++")
	fmt.Println()
	list.AddFromHead(1)
	list.AddFromHead(2)
	list.AddFromHead(3)
	print(list)

	fmt.Println("++++++++3、向双向链表尾部添加元素++++++++")
	fmt.Println()
	list.Append("Golang")
	list.Append("PHP")
	list.Append("Java")
	print(list)

	fmt.Println("++++++++4、向双向链表的指定位置插入结点++++++++")
	fmt.Println("++++++++(1)向双向链表的【第一个】位置插入结点++++++++")
	list.Insert(1, "First")
	print(list)

	fmt.Println("++++++++(2)向双向链表的【第三个】位置插入结点++++++++")
	list.Insert(3, "Third")
	print(list)

	fmt.Println("++++++++(3)向双向链表的【最后】位置插入结点++++++++")
	list.Insert(list.Length()+1, "Last")
	print(list)

	fmt.Println("++++++++5、删除双向链表的头结点++++++++")
	fmt.Println()
	list.RemoveHeadNde()
	print(list)

	fmt.Println("++++++++6、删除双向链表的尾结点++++++++")
	fmt.Println()
	list.RemoveLastNode()
	print(list)

	fmt.Println("++++++++7、删除双向表中指定值的结点++++++++")
	list.Remove(3)//这里的类型要和你插入时的类型一致（弱类型语言写习惯了，很容易忘了）
	print(list)

	fmt.Println("++++++++8、删除双向表中指定位置的结点++++++++")
	fmt.Println("++++++++(1)删除双向链表的【第一个】位置结点++++++++")
	list.RemovePosition(1)
	print(list)

	fmt.Println("++++++++(2)删除双向链表的【第三个】位置结点++++++++")
	list.RemovePosition(3)
	print(list)

	fmt.Println("++++++++(3)删除双向链表的【最后】位置结点++++++++")
	list.RemovePosition(list.Length())
	print(list)

	fmt.Println("++++++++9、查询双向链表中是否包含某一个结点++++++++")
	fmt.Println()
	fmt.Printf("是否存在Golang结点：%v\n", list.Contain(2))
}
