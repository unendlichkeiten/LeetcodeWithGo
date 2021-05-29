package _167_two_sum

func twoSumByMap(numbers []int, target int) []int {
	numbersMap := make(map[int]int)

	n := 0
	for i, v := range numbers {
		n = target - v
		if _, ok := numbersMap[n]; ok {
			return []int{numbersMap[n] + 1, i + 1}
		} else {
			numbersMap[v] = i
		}
	}

	return nil
}

// condition: ordered sequence
func twoSumByIndex(numbers []int, target int) []int {
	left, right := 1, len(numbers)

	for left < right {
		if numbers[left-1]+numbers[right-1] > target {
			right--
		} else if numbers[left-1]+numbers[right-1] < target {
			left++
		} else {
			return []int{left, right}
		}
	}

	return nil
}
