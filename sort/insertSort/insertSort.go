package insertSort

import "fmt"

func InsertSort(arr []int)  {
	if len(arr) <= 0 {
		fmt.Println("待排数据不合法")
	}
	n := len(arr)

	for i := 1; i < n; i++ {//i是待排区的元素
		value := arr[i]
		j := i-1
		for ; j >= 0; j-- { //j遍历的是已排区的每一个元素
			if arr[j] > value {
				arr[j+1] = arr[j] //如果满足条件，将前一个值赋给后边这个
			} else {
				break
			}
		}
		arr[j+1] = value
	}

	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}
}
