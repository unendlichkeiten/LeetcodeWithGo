package _605_can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	left, right, count := -1, len(flowerbed), 0
	for i := 0; i < right; i++ {
		if flowerbed[i] == 1 {
			if left < 0 {
				count += i / 2
			} else {
				count += (i - left - 2) / 2
			}
			left = i
		}
	}

	if left < 0 {
		count += (right + 1) / 2
	} else {
		count += (right - left - 1) / 2
	}

	return count >= n
}
