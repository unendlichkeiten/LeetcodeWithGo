package _000_common_sort_algorithm

import (
	"math"
)

type elem struct {
	num  int
	next *elem
}

func BucketSort(nums []int, bucketN int) {
	min, max := math.MaxInt32, math.MinInt32

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
	i, interval := 0, (max-min)/(bucketN-1)
	if bucketN > 1 {
		interval = (max - min) / (bucketN - 1)
	} else {
		interval = max - min
	}
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
				next: p,
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
