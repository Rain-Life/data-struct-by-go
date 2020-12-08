package main

import (
	"data-struct-by-go/hashTable/practice"
	"fmt"
)

func main() {
	//对于10w个url，实战场景中可以通过读取文件来获取，假设url平均长度为50，
	//10w个url也就不到5M，完全可以读取到内存中进行处理
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
