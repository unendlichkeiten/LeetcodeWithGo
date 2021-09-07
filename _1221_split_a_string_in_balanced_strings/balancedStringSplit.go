package _1221_split_a_string_in_balanced_strings

func balancedStringSplit(s string) int {
	/*
	   1.按照顺序遍历字符串 s，numR, numL 分别表示遍历过程中两个字符出现的个数
	   2.若 numR == numL 且 numR > 0 && numL > 0, 平衡字符串个数+1
	*/

	numR, numL, balancedS := 0, 0, 0
	for _, c := range s {
		if c == 'R' {
			numR++
		}

		if c == 'L' {
			numL++
		}

		if numR > 0 && numR == numL {
			balancedS++
		}
	}

	return balancedS
}
