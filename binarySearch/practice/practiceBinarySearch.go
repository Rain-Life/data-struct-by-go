package practice

import "fmt"

//查找第一个值等于给定值的元素
func BinarySearchFirst(arr []int, value int) int {
	if len(arr) <= 0 {
		fmt.Println("数组是空的")
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high-low) >> 1)
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || (arr[mid-1] != value) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, value int) int {
	if len(arr) <= 0 {
		fmt.Println("数组是空的")
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high-low) >> 1)
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == (len(arr) - 1) || (arr[mid+1] != value) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchFirstBigger(arr []int, value int) int {
	if len(arr) <= 0 {
		fmt.Println("数组是空的")
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high-low) >> 1)
		if arr[mid] >= value {
			if mid == 0 || arr[mid -1] < value {
				return mid
			} else {
				high = mid -1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}

//查找最后小于等于给定值的元素
func BinarySearchLastSmaller(arr []int, value int) int {
	if len(arr) <= 0 {
		fmt.Println("数组是空的")
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high-low) >> 1)
		if arr[mid] <= value {
			if (mid == (len(arr) - 1)) || arr[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}

	return -1
}