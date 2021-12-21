package BinarySearch

import (
	"fmt"
)

//二分查找，非递归实现
func BinarySearch(arr []int, value int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > value {
			high = mid-1
		} else if arr[mid] < value {
			low = low + 1
		} else {
			return mid
		}
	}

	return -1
}

//二分查找：递归实现
func RecursionBinarySearch(arr []int, low, high, value int) int {
	if low < high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == value {
		return mid
	} else if arr[mid] > value {
		return RecursionBinarySearch(arr, low, mid - 1, value)
	} else {
		return RecursionBinarySearch(arr, mid +	1, high, value)
	}
}


//如果待查的数据中有重复的值，则取第一个等于待查的值
func BinarySearchFirst(arr []int, low, high, value int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] < value {
		return BinarySearchFirst(arr, mid + 1, high, value)
	} else if arr[mid] > value {
		return BinarySearchFirst(arr, low, mid - 1, value)
	} else {
		if mid == 0 || (arr[mid-1]) != value {
			return mid
		} else {
			return BinarySearchFirst(arr, low, mid - 1, value)
		}
	}
}


//如果待查的数据中有重复的值，则取最后一个等于待查的值
func BinarySearchLast(arr []int, low, high, value int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] < value {
		return BinarySearchLast(arr, mid + 1, high, value)
	} else if arr[mid] > value {
		return BinarySearchLast(arr, low, mid - 1, value)
	} else {
		if (mid == (len(arr)-1)) || (arr[mid+1] != value) {
			return mid
		} else {
			return BinarySearchLast(arr, mid+1, high, value)
		}
	}
}


//寻找第一个大于等于待查找的值（同上）


//寻找最后一个大=小于等于待查找的值（同上）





//先看一个简单的数组移动的问题，假设将数组向左或向右移动n位，打印出移动后的数组（数组中的数字不一定是有序，这里只是以有序数组为例）
// operateType(l左移， r右移)

//方法一：利用数组的切割，比如1、2、3、4、5、6、7，右移4位，就是4、5、6、7、1、2、3，那就相当于把前3个切割出来，放到后边
func MoveArray(arr []int, n int, operateType rune) {
	n = n % len(arr)
	if n == 0 {
		fmt.Println(arr)
		return
	}
	if operateType == 'r' { //右移n位，就等于左移len - n
		n = len(arr) - n
	}

	arr1 := arr[n:]
	arr2 := arr[:n]
	fmt.Println(arr1, arr2)
	newArr := append(arr1, arr2...)
	fmt.Println(newArr)
}

//方法二：假设将一个数组向坐移动3位，可以通过以下几步来完成  变成这样的数组[4, 5, 6, 7, 1, 2, 3]（原数组是[1, 2, 3, 4, 5, 6, 7]）
/**
 * 1.首先将前三个数字进行逆转，就变成了[3, 2, 1, 4, 5, 6, 7]
 * 2.再将后面四个数组进行逆转，就变成了[3, 2, 1, 7, 6, 5, 4]
 * 3.最后我们将数组整体进行逆转，就变成了[4, 5, 6, 7, 1, 2, 3]
 */
//说明：像方法一一样，向右移动n位，就等于向左移动len - n
func MoveArrayWay2(arr []int, n int, operateType rune) {
	n = n % len(arr)
	if n == 0 {
		fmt.Println(arr)
		return
	}
	if operateType == 'r' { //右移n位，就等于左移len - n
		n = len(arr) - n
	}

	i := 0
	j:=n-1
	for i<j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)

	i = n
	j = len(arr) - 1
	for i<j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)

	i = 0
	j = len(arr) - 1
	for i<j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)
}




//整数数组 nums 按升序排列，数组中的值 互不相同(leetcode 33)
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
//使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2]
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-er-fen-92jot/

func FindTarget(arr []int, target int) int {
	left, right := 0, len(arr)-1

	//将旋转后的数组看做两个严格升序的数组
	for left <= right {
		mid := left + (right-left) >> 1

		if arr[mid] == target {
			return mid
		}

		//落在同一数组的情况，同时落在数组1 或 数组2
		if arr[mid] >= arr[left] {
			//target 落在 left 和 mid 之间，则移动我们的right，完全有序的一个区间内查找
			if arr[mid] > target && target >= left {
				right = mid-1
			} else if arr[mid] < target || target < arr[left] {//target 落在right和 mid 之间，有可能在数组1， 也有可能在数组2(你也可以直接else，我这里是为了方便理解)
				left = mid+1
			}
		}
		//不落在同一数组的情况，left 在数组1， mid 落在 数组2
		if arr[mid] < arr[left] {
			// 有序的一段区间，target 在 mid 和 right 之间
			if target > arr[mid] && target <= arr[right] {
				left = mid+1
			} else if target < arr[mid] || target > arr[right] {// 两种情况，target 在left 和 mid 之间
				right = mid-1
			}
		}
	}

	return -1
}





