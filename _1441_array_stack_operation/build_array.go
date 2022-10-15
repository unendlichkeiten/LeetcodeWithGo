package _1441_array_stack_operation

func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	for index, val := range target {
		if index > 0 && target[index]-target[index-1] == 1 {
			res = append(res, "Push")
			continue
		}

		start := 0
		if index == 0 {
			start = index + 1
		} else {
			start = target[index-1] + 1
		}
		for n := start; n < val; n++ {
			res = append(res, "Push")
			res = append(res, "Pop")
		}

		res = append(res, "Push")

	}

	return res
}
