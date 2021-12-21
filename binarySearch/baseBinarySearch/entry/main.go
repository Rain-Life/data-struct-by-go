package main

import (
	"data-struct-by-go/binarySearch/baseBinarySearch"
)

func main() {
	var arr = []int{1,2,3,4,5,6,7}
	//fmt.Println(BinarySearch.BinarySearchLast(arr, 0, len(arr)-1, 5))
	//BinarySearch.MoveArray(arr, 10, 'r')
	BinarySearch.MoveArrayWay2(arr, 3, 'l')
}
