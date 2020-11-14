package main

import (
	"data-struct-by-go/sort/quickSort"
	"fmt"
)

func main() {
	arr := []int{6,1,3,5,7,2,4,9,11,8}
	//fmt.Println("排序之前")
	//fmt.Println(arr)
	//quickSort.QuickSort(arr, 0, len(arr)-1)
	//fmt.Println("排序之后")
	//fmt.Println(arr)

	topK := quickSort.FindTopK(arr, 3)
	fmt.Println("第三大的元素是")
	fmt.Println(topK)
}
