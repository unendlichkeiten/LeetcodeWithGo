package leetCode

import "errors"

// 题目一：找出重复的数字
// 方法二：hash表法
func ReplicaValue_Hash(array []int)(int, error){
	arrMap := make(map[int]int)

	for _, value := range array {
		if _, ok := arrMap[value]; ok {
			return value, nil
		} else {
			arrMap[value] = 1
		}
	}

	return -1, errors.New("no value replica")
}

// 题目一：找出重复的数字
// 方法二：扫描法
func ReplicaValue_Scan(array []int)(int, error){
	for key, value := range array {
		compare:
		if value == key {
			continue
		} else {
			if array[value] == value {
				return value, nil
			} else {
				array[value], value = value, array[value]
				goto compare
			}
		}
	}

	return -1, errors.New("no value replica")
}

// 获取枢轴
func partition(num []int, left, right int) int {
	pivot := num[left]
	for ; left < right; {
		for ; left < right && num[right] >= pivot; right-- {}
		num[left] = num[right]
		for ; left < right && num[left] <= pivot; left++ {}
		num[right] = num[left]
	}

	num[left] = pivot
	return left
}

// 快排
func QuickSort2(num []int, left, right int) {
	if left < right {
		pivot := partition(num, left, right)
		QuickSort2(num, left, pivot)
		QuickSort2(num, pivot+1, right)
	}
}

// 题目二：不修改数组找出重复数字
// 方法一：空间换时间
func RepicaValue_Mem(array []int) (int, error) {
	tmpArray := make([]int, 0)
	for _, value := range array {
		tmpArray = append(tmpArray, value)
	}

	QuickSort2(tmpArray, 0, len(array)-1)

	for i, value := range tmpArray{
		if i != 0 && value == tmpArray[i-1] {
			return value, nil
		}
	}
	return -1, errors.New("not exists replica value")
}

// 题目二：不修改数组找出重复数字
// 方法二：二分法
func ReplicaValue_Div(array []int, left, right int) (int, error) {
	for left < right {
		m, n, i := (left+right)/2, (right-left)/2+1, 0
		for key, value := range array {
			if value <= m && key > 0 {
				i++
			}
		}

		if i > n {
			right = m
		} else {
			left = m + 1
		}
	}

	return left, nil
}

