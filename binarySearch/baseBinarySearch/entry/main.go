package main

import (
	"data-struct-by-go/binarySearch/baseBinarySearch"
	"fmt"
)

func main() {
	arr := []int{2,3,4,5,6,7,8,9}
	position := baseBinarySearch.BinarySearch(arr, 2)
	fmt.Println(position)

	position1 := baseBinarySearch.RecursionBinarySearch(arr, 0, len(arr)-1, 10)
	fmt.Println(position1)
}
