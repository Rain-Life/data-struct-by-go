package quickSort

func QuickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	position := partition(arr, start, end)
	QuickSort(arr, start, position-1)
	QuickSort(arr, position+1, end)
}


func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start
	for j:=start;j < end;j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j],arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]

	return i
}


