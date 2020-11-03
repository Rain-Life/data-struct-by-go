package bubleSort

import "fmt"

func BubbleSort(arr []int) {
	flag := false
	n := len(arr)
	for i := 0; i < n; i++ {
		flag = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}
}