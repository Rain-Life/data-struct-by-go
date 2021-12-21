package DynamicProgramming


// 0-1 背包问题
//weight：物品重量 n：物品数量 w：背包可承载的重量
func Knapsack(weight []int, n, w int) int {
	states := [5][10]bool{}

	states[0][0] = true //对第一个物品进行特殊处理，哨兵模式
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i := 1;i < n; i++ { // 动态规划，状态转移
		for j := 0; j <= w; j++ { //不把第i个物品放入背包
			if states[i-1][j] == true { // 如果i的上一个物品放入背包了，直接把它的状态拷贝下来就行（因为i不放）
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= w - weight[i]; j++ { // 把第i个物品放入背包
			if states[i-1][j] == true {
				states[i][j + weight[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[n-1][i] == true {
			return i
		}
	}

	return 0
}

// 0-1背包问题，对空间复杂度进行优化

func Knapsack2(items []int, n, w int) int {
	states := [10]bool{}

	states[0] = true
	if items[0] <= w {
		states[items[0]] = true
	}

	for i := 0; i < n; i++ {
		for j := w-items[i]; j >= 0; j-- { // 将第i个物品放入背包
			if states[j] == true {
				states[j + items[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[i] ==  true {
			return i
		}
	}

	return 0
}