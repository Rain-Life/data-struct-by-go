package baseBinarySearch

import "fmt"

//二分查找的非递归实现
func BinarySearch(arr []int, value int) int {
	if len(arr) < 0 {
		fmt.Println("空数组")
		return -1
	}
	low := 0;
	high := len(arr)-1
	for low <= high {
		mid := low + ((high-low) >> 1)
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			return mid
		}
	}

	fmt.Println("数组中不存在要找的值")
	return -1
}

//二分查找的递归实现
func RecursionBinarySearch(arr []int, low int, high int, value int) int {
	if len(arr) <= 0 {
		fmt.Println("数组是空的")
		return -1
	}
	if low > high { //注意这里没有=，low==high时还不能判定数组中不存在value，画一下就明白了
		return -1
	}

	mid := low + ((high-low) >> 1)
	if arr[mid] == value {
		return mid
	} else if arr[mid] > value {
		return RecursionBinarySearch(arr, low, mid-1, value)
	} else {
		return RecursionBinarySearch(arr, mid + 1, high, value)
	}
}
