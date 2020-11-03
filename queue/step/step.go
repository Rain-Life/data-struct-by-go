package step

import "fmt"

//假如这里有 n 个台阶，每次你可以跨 1 个台阶或者 2 个台阶，请问走这 n 个台阶有多少种走法？如果有 7 个台阶，
//你可以 2，2，2，1 这样子上去，也可以 1，2，1，1，2 这样子上去，总之走法有很多，那如何用编程求得总共有多少种走法呢？

var depth = 0

var mapList = map[int]int{}

func Step(n int) int {
	depth++
	if depth > 100 {
		panic("递归次数太多")
	}
	if n == 1 {
		return 1
	} else if n==2 {
		return 2
	}

	if mapList[n] != 0 {
		return mapList[n]
	}
	res := Step(n-1)+Step(n-2)
	mapList[n] = res
	fmt.Printf("step(%d) = %d", n, mapList[n])
	fmt.Println()
	return res

}
