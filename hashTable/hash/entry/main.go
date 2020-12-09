package main

import "data-struct-by-go/hashTable/hash"

func main() {
	var hashTable hash.HashTable
	for i:=1; i < 1000;i++ {
		data := hash.Data{Value: i}
		hashTable.InsertHashTable(&data)
	}
	hashTable.ShowHashTable()
}
