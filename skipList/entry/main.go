package main

import (
	"data-struct-by-go/hashTable/practice"
	"fmt"
)

func main() {
	urls := []string{
		"http://www.baidu.com/1",
		"http://www.baidu.com/2",
		"http://www.baidu.com/3",
		"http://www.baidu.com/2",
		"http://www.baidu.com/3",
		"http://www.baidu.com/1",
		"http://www.baidu.com/1",
		"http://www.baidu.com/4",
		"http://www.baidu.com/1",
		"http://www.baidu.com/9",
		"http://www.baidu.com/3",
		"http://www.baidu.com/1",
		"http://www.baidu.com/1",
		"http://www.baidu.com/1",
	}
	hashTable := practice.MakeMap(urls)
	fmt.Println(hashTable)
	urlStruct := practice.MapToStruct(hashTable)
	fmt.Println(urlStruct)
	urlStruct.QuickSort(0, len(urlStruct)-1)
	fmt.Println(urlStruct)
}
