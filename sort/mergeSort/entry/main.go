package main

import (
	"data-struct-by-go/sort/mergeSort"
	"fmt"
)

func main() {
	arr := []int{6, 4, 5, 1, 8, 7, 2, 3}
	fmt.Println("排序之前")
	fmt.Println(arr)
	resArr := mergeSort.MergeSort(arr)
	fmt.Println("排序之后")
	fmt.Println(resArr)
}
