package _103_minimum_coin

func coinChange(coins []int, amount int) int {
	// 构建状态转移方程
	// F(X) = MIN(F(X-coins[0]), F(X-coins[1]), ... F(X-coins[len(coins)-1])) + 1
	if amount == 0 {
		return 0
	}

	minCoins := make(map[int]int)
	for _, coin := range coins {
		if amount > coin {
			minCoins[amount] = min(coinChange(coins, amount-coin), minCoins[amount])
		}
	}

	return minCoins[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
