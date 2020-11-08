package main

import (
	"data-struct-by-go/binarySearch/practice"
	"fmt"
)

func main() {
	arr := []int{1,3,8,8,8,10,13}
	//查找第一个值等于给定值的元素
	positionFirst := practice.BinarySearchFirst(arr, 8)
	fmt.Println(positionFirst)

	//查找最后一个值等于给定值的元素
	positionLast := practice.BinarySearchLast(arr, 8)
	fmt.Println(positionLast)

	//查找第一个大于等于给定值的元素
	positionBig := practice.BinarySearchFirstBigger(arr, 8)
	fmt.Println(positionBig)

	//查找最后小于等于给定值的元素
	positionSmall := practice.BinarySearchLastSmaller(arr, 8)
	fmt.Println(positionSmall)
}
