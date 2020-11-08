package baseBinarySearch

import "fmt"

func BinarySearch(arr []int, value int) int {
	if len(arr) < 0 {
		fmt.Println("空数组")
		return 0
	}
	low := 0;
	height := len(arr)-1
	for low <= height {
		mid := (low+height)/2
		if arr[mid] > value {
			height = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			return mid
		}
	}

	fmt.Println("数组中不存在要找的值")
	return 0
}
