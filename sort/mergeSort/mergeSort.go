package mergeSort

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 0 {
		fmt.Println("参数不合法")
		return nil
	}

	//递归终止条件，当拆分后的数组长度少于两个元素的时候
	if len(arr) < 2 {
		return arr
	}

	midIndex := len(arr) / 2
	leftArr := MergeSort(arr[0:midIndex])
	rightArr := MergeSort(arr[midIndex:])
	result := merge(leftArr, rightArr)

	return result
}

func merge(leftArr, rightArr []int) []int {
	var mergeRes []int
	leftIndex, rightIndex := 0, 0
	leftLength, rightLength := len(leftArr), len(rightArr)

	for leftIndex < leftLength && rightIndex < rightLength {
		if leftArr[leftIndex] > rightArr[rightIndex] {
			mergeRes = append(mergeRes, rightArr[rightIndex])
			rightIndex++
		} else {
			mergeRes = append(mergeRes, leftArr[leftIndex])
			leftIndex++
		}
	}

	if leftIndex == leftLength{
		for rightIndex < rightLength {
			mergeRes = append(mergeRes, rightArr[rightIndex])
			rightIndex++
		}
	}
	if rightIndex == rightLength {
		for leftIndex < leftLength {
			mergeRes = append(mergeRes, leftArr[leftIndex])
			leftIndex++
		}
	}

	return mergeRes
}

