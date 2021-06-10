package _000_common_sort_algorithm

import (
	"math"
)

type elem struct {
	num  int
	next *elem
}

func BucketSort(nums []int) {
	min, max, numsLen := math.MaxInt32, math.MinInt32, len(nums)

	// find get min and max value
	for _, n := range nums {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	// figure out bucket capacities and put num into bucket
	i, interval := 0, (max-min)/numsLen
	bucket := make(map[int]*elem)
	var pre *elem = nil
	for _, n := range nums {
		i = (n - min) / interval
		p := bucket[i]
		for p != nil && n > p.num {
			pre = p
			p = p.next
		}

		if pre == nil {
			bucket[i] = &elem{
				num:  n,
				next: nil,
			}
		} else {
			pre.next = &elem{
				num:  n,
				next: p,
			}
		}
	}

	// travel bucket value into source slice
	index := 0
	for i := 0; i < len(bucket); i++ {
		p := bucket[i]
		for p != nil {
			nums[index] = p.num
			index++
			p = p.next
		}
	}
}
