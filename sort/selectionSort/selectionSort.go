package selectionSort

import "fmt"

func SelectionSort(arr []int) {
	n := len(arr)
	if n <= 0 {
		fmt.Println("待排数据不合法")
	}

	for i := 0; i < n - 1; i++ {
		for j := i+1; j < n ; j++ {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j], arr[i]
			}
		}
	}

	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}
}

