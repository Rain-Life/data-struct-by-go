package main

import (
	"data-struct-by-go/sort/quickSort"
	"fmt"
)

func main() {
	arr := []int{6, 4, 5, 1, 8, 7, 2, 3}
	fmt.Println("排序之前")
	fmt.Println(arr)
	quickSort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("排序之后")
	fmt.Println(arr)
}
