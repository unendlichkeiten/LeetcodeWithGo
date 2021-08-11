package _103_minimum_coin

import "math"

func coinChange(coins []int, amount int) int {
	// 构建状态转移方程
	// F(X) = MIN(F(X-coins[0]), F(X-coins[1]), ... F(X-coins[len(coins)-1])) + 1
	if amount == 0 {
		return 0
	}

	minCoins := make([]int, amount)
	for i := 1; i <= amount; i++ {
		minCoins[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] &&
				minCoins[i-coins[j]] != math.MaxInt32 &&
				minCoins[i-coins[j]]+1 < minCoins[i] {
				minCoins[i] = minCoins[i-coins[j]] + 1
			}
		}
	}

	if minCoins[amount] == math.MaxInt32 {
		return -1
	}

	return minCoins[amount]
}
