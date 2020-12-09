package hash

import (
	"fmt"
	"math/rand"
)

//实现hash表，并通过链表法处理散列冲突
const SIZE = 100 //散列表大小

//定义实际存储数据的结构体
type Data struct {
	Value int
	Next *Data
}

//DataList为每一个槽对应的值组成的链表
type DataLinked struct {
	Head *Data
}

//定义散列表
type HashTable struct {
	LinkedArr [SIZE]DataLinked
}

//实现哈希函数（这里简单实现以下）
func (hashTable *HashTable) HashFun(key int32) int32 {
	return key % SIZE
}

//向散列表中插入一条数据
func (hashTable *HashTable) InsertHashTable(data *Data) {
	//随机生成一个键值，然后生成哈希值
	//rand.Int31() int32  [0, MaxInt32]
	hashKey := hashTable.HashFun(rand.Int31())
	//将数据插入到对应槽后边的链表中(真正的插入数据是在这一步)
	hashTable.LinkedArr[hashKey].InsertDataLinked(data)
}

//向散列表的槽中插入具体数据
func (dataLinked *DataLinked) InsertDataLinked(data *Data) {
	currentNode := dataLinked.Head

	//判断是否存在散列冲突（即该槽中是否已经有数据）
	if dataLinked.Head == nil {//不存在散列冲突
		dataLinked.Head = data
		return
	}

	//存在散列冲突
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = data
}

//根据键值及Value查询对应的数据
func (hashTable *HashTable) FindHashTableByKey(key int32, value int) *Data {
	hashKey := hashTable.HashFun(key)
	return hashTable.LinkedArr[hashKey].FindLinkedByValue(value)
}

//在对应槽后边的链表中查找指定值的节点
func (dataLinked *DataLinked) FindLinkedByValue(value int) *Data {
	if dataLinked.Head == nil {
		return nil
	}

	currentNode := dataLinked.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}

		currentNode = currentNode.Next
	}

	return nil
}

//展示某一个槽的值（链表）
func (dataLinked *DataLinked) ShowDataLinked(hashKey int32) {
	if dataLinked.Head == nil {
		fmt.Printf("槽%d 为空\n", hashKey)
		return
	}

	currentNode := dataLinked.Head
	for currentNode != nil {
		fmt.Printf("槽%d:Value:%d ->", hashKey, currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

//打印散列表
func (hashTable *HashTable) ShowHashTable() {
	for i:=0; i < len(hashTable.LinkedArr); i++ {
		hashTable.LinkedArr[i].ShowDataLinked(int32(i))
	}
}
