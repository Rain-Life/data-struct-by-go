package main

import (
	"data-struct-by-go/queue/step"
	"fmt"
)

func main() {
	n := 12
	count := step.Step(n)
	fmt.Printf("%d个台阶一共有%d中走法", n, count)
}
