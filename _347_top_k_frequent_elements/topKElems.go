package _347_top_k_frequent_elements

type elem struct {
	num   int
	count int
	next  *elem
}

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		// record each element frequent
		numsMap[v]++
	}

	var pre *elem = nil
	var cur *elem = nil
	var root *elem = nil

	for k, c := range numsMap {
		if root == nil {
			root = &elem{
				num:   k,
				count: c,
				next:  nil,
			}
		} else {
			for cur != nil && c <= cur.count {
				pre = cur
				cur = cur.next
			}

			if pre == nil {
				root = &elem{
					num:   k,
					count: c,
					next:  cur,
				}
			} else {
				pre.next = &elem{
					num:   k,
					count: c,
					next:  cur,
				}
			}
		}
		cur = root
		pre = nil
	}

	i, topKs, cur := 0, make([]int, 0), root
	for i < k {
		topKs = append(topKs, cur.num)
		i++
		cur = cur.next
	}

	return topKs
}
