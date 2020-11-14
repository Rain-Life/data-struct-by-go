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

//“求一个数的平方根”？要求精确到小数点后 6 位
func SolutionSqrt(num float64, precision float64) float64 {
	if num < 0 {
		fmt.Println("参数不合法")
		return -1.000
	}

	var low,high float64
	if num > 1 {
		low = 1
		high = num
	} else {
		low = num
		high = 1
	}

	return Sqrt(num, low, high, precision)
}

func Sqrt(num float64, low float64, high float64, precision float64) int {
	mid := (low+high) / 2

	if num - (mid * mid) < 0 {
		return Sqrt(num, low, (mid+high)/2, precision)
	} else {
		if num - (mid * mid) > precision {
			return Sqrt(num, (low + mid)/2, high, precision)
		}

		return mid.(int)
	}
}