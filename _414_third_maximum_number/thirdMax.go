package _414_third_maximum_number

import "math"

/*
** 解题思路：
** 1.声明三个局部变量 max1, max2, max3, 赋值 math.MinInt64
** 2.遍历数组，一次和三个变量比较大小，若该值等于其中某个变量值直接跳过
** 3.返回结果，若 max3 存在返回该值，若不存在返回 max1
 */
func thirdMax(nums []int) int {
	// 分别表示三个最大的数，对于重复出现的数值只记录一个
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 && num < max1 {
			max2, max3 = num, max2
		} else if num > max3 && max2 > max3 && num < max2 {
			max3 = num
		}
	}

	if max3 == math.MinInt64 {
		return max1
	}

	return max3
}
