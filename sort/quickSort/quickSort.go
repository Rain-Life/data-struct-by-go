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


