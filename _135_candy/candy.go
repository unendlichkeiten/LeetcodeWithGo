package _135_candy

func candy(ratings []int) int {
	candyn := make([]int, len(ratings))
	for i := range candyn {
		candyn[i] = 1
	}

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			candyn[i+1] = candyn[i] + 1
		}
	}

	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candyn[i-1] <= candyn[i] {
			candyn[i-1] = candyn[i] + 1
		}
	}

	n := 0
	for i := range candyn {
		n += candyn[i]
	}

	return n
}
