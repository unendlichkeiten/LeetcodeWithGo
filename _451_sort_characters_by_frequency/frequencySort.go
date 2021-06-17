package _451_sort_characters_by_frequency

type elem struct {
	ch    int32
	count int
	next  *elem
}

func frequencySort(s string) string {
	// record each char frequency
	strMap := make(map[int32]int)
	for _, ch := range s {
		strMap[ch]++
	}

	var pre *elem = nil
	var cur *elem = nil
	var root *elem = nil

	for ch, c := range strMap {
		if root == nil {
			root = &elem{
				ch:    ch,
				count: c,
				next:  nil,
			}
		} else {
			// find insert pos
			for cur != nil && c <= cur.count {
				pre = cur
				cur = cur.next
			}

			// insert elem
			if pre == nil {
				root = &elem{
					ch:    ch,
					count: c,
					next:  cur,
				}
			} else {
				pre.next = &elem{
					ch:    ch,
					count: c,
					next:  cur,
				}
			}
		}
		cur = root
		pre = nil
	}

	sortStr := make([]int32, 0)
	for cur != nil {
		for i := 0; i < cur.count; i++ {
			sortStr = append(sortStr, cur.ch)
		}
		cur = cur.next
	}

	return string(sortStr)
}
