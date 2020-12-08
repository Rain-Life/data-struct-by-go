package _46

import "fmt"

var result [][]int
func Permute(nums []int) {
	doPermute(nums, len(nums), len(nums))
}

func doPermute(arr []int, n int, k int) {
	if k == 1 {
		//tmp := make([]int, len(arr))
		//copy(tmp, arr)
		//result = append(result, tmp)
		fmt.Println(arr)
	}

	for i := 0;i < k; i++ {
		arr[i], arr[k-1] = arr[k-1], arr[i]
		doPermute(arr, n, k-1)
		arr[i], arr[k-1] = arr[k-1], arr[i]
	}
}

