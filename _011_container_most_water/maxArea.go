package _011_container_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	vmax := (right - left) * min(height[left], height[right])
	for left < right && left < len(height) && right > 0 {
		if v := (right - 1 - left) * min(height[left], height[right-1]); v > vmax {
			vmax = v
			right--
		}

		if v := (right - left - 1) * min(height[left-1], height[right]); v > vmax {
			vmax = v
			left++
		}

	}

	return vmax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
