package _000_common_sort_algorithm

import "math"

func shellSort(nums []int) []int {
	lenNums := len(nums)
	for dk := lenNums / 2; dk > 0; dk /= 2 {
		// 这里会不断切换子表，比较难搞
		for i := dk; i < lenNums; i++ {
			if nums[i] < nums[i-dk] {
				// nums[0] 属于哨兵，序列起始位置为1
				tmp, j := nums[i], math.MinInt32
				for j = i - dk; j >= 0 && tmp < nums[j]; j -= dk {
					nums[j+dk] = nums[j]
				}
				nums[j+dk] = tmp
			}
		}
	}

	return nums
}

func shellSort2(nums []int) []int {
	lenNums := len(nums)
	for dk := lenNums / 2; dk > 0; dk /= 2 {
		// 这里不切换子表
		for j := 0; j < dk; j++ {
			for i := dk + j; i < lenNums; i += dk {
				if nums[i] < nums[i-dk] {
					// nums[0] 属于哨兵，序列起始位置为1
					tmp, k := nums[i], math.MaxInt32
					for k = i - dk; k > 0 && tmp < nums[k]; k -= dk {
						nums[k+dk] = nums[k]
					}
					nums[k+dk] = tmp
				}
			}
		}
	}

	return nums
}
