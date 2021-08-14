package _1582_count_unhappy_friends

import (
	"math"
)

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	// 初始化亲密度矩阵
	relationship := make([][]int, n)
	for i := 0; i < n; i++ {
		relationship[i] = make([]int, n)
		for j := 0; j < n; j++ {
			relationship[i][j] = math.MaxInt32
		}
	}

	// 记录各个朋友亲密程度 i 代表各个朋友的编号
	for i := 0; i < n; i++ {
		for k, v := range preferences[i] {
			relationship[i][v] = k
		}
	}

	// 记录配对朋友
	forcePairs := make(map[int]int)
	for _, pair := range pairs {
		forcePairs[pair[0]] = pair[1]
		forcePairs[pair[1]] = pair[0]
	}

	var unhappyMembers int
	// 遍历每一组成员，找到他/她爱的人
	for _, pair := range pairs {
		relation := []int{
			relationship[pair[0]][pair[1]], // relationXToY
			relationship[pair[1]][pair[0]], // relationYToX
		}

		// 判断组内的每个成员是否幸福
		for XY, memberX := range pair {
			for memberU, relationXU := range relationship[memberX] {
				flag := false
				if relationXU < relation[XY] { // 说明 memberX 没有和爱的人在一起
					relationUX := relationship[memberU][memberX]

					// 查找 memberX 爱人 memberU 的爱人
					for memberV, relationUV := range relationship[memberU] {
						if forcePairs[memberU] == memberV && relationUX < relationUV {
							// 说明 memberX 的爱人也没有和她/他更爱的人在一起，当然就不幸福咯
							unhappyMembers++
							flag = true
							break
						}
						// 这里当然说明 memberX 的爱的人的对象，比她/他更好。
						// memberX 当然成全她/他咯，哈哈
					}

					// 确定 memberX 不幸福直接找找下一个不幸福的人
					if flag {
						break
					}
				}
			}
		}
	}

	return unhappyMembers
}
