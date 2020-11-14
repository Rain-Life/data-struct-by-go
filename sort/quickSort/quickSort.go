package quickSort

func QuickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := partition(arr, start, end)
	QuickSort(arr, start, pivot-1)
	QuickSort(arr, pivot+1, end)
}


func partition(arr []int, start int, end int) int {
	pivotValue := arr[end]
	i := start
	for j:=start;j < end;j++ {
		if arr[j] < pivotValue {
			arr[i], arr[j] = arr[j],arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]

	return i
}

// 在 O(n) 的时间复杂度内查找一个无序数组中的第 K 大元素？

func FindTopK(arr []int, K int) int {
	return TopK(arr, 0, len(arr)-1, K)
}

func TopK(arr []int, start int, end int, K int) int {
	//if start >= end { 这里是个坑，不能加这个限制条件
	//	return -1
	//}
	pivot := partition(arr, start, end)

	if pivot+1 == K {
		return arr[pivot]
	} else if pivot+1 < K {
		return TopK(arr, pivot+1, end, K)
	} else {
		return TopK(arr, start, pivot-1, K)
	}
}
