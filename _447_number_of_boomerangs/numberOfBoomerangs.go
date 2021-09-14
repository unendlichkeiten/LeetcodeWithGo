package _447_number_of_boomerangs

func NumberOfBoomerangs(points [][]int) int {
	var sums int
	for _, p := range points {
		cnts := make(map[int]int)
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnts[dis]++
		}

		for _, cnt := range cnts {
			sums += cnt * (cnt - 1)
		}
	}

	return sums

}
