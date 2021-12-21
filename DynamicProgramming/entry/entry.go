package main

import (
	"data-struct-by-go/DynamicProgramming"
	"fmt"
)

func main() {
	weight := []int{2,2,4,6,3}
	fmt.Println(DynamicProgramming.Knapsack(weight, len(weight), 9))
}
