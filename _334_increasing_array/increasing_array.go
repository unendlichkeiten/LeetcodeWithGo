package _334_increasing_array

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	numsArr := make([][]int, 0)
	for _, num := range nums {
		greater, smaller, tmp := math.MaxInt32, math.MaxInt32, 0
		for i, arr := range numsArr {
			if len(arr) > 1 && num < arr[len(arr)-1] &&
				num-arr[len(arr)-2] > 0 && num-arr[len(arr)-2] < smaller {
				smaller = i
			}

			if len(arr) > 1 && num > arr[len(arr)-1] {
				numsArr[i] = append(numsArr[i], num)
				break
			} else if len(arr) > 0 && num > arr[len(arr)-1] && num-arr[len(arr)-1] < greater {
				greater = i
			}
		}

		if smaller < math.MaxInt32 {
			tmp = numsArr[smaller][len(numsArr[smaller])-1]
			numsArr[smaller][len(numsArr[smaller])-1] = num
			num = tmp
		}

		if greater < math.MaxInt32 {
			numsArr[greater] = append(numsArr[greater], num)
			continue
		}

		numsArr = append(numsArr, []int{num})
	}

	for _, arr := range numsArr {
		if len(arr) >= 3 {
			return true
		}
	}

	return false
}
