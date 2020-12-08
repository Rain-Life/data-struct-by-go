package main

import "data-struct-by-go/hashTable/practice"

func main() {
	urls := []string{
		"http://www.baidu.com/1",
		"http://www.baidu.com/1",
		"http://www.baidu.com/2",
		"http://www.baidu.com/2",
		"http://www.baidu.com/1",
		"http://www.baidu.com/3",
		"http://www.baidu.com/4",
		"http://www.baidu.com/5",
		"http://www.baidu.com/4",
		"http://www.baidu.com/1",
		"http://www.baidu.com/2",
		"http://www.baidu.com/1",
	}
	practice.TenWUrlSort(urls)
}
