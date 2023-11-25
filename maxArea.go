package main

func maxVal(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxArea(heights []int) int {
	i, j := 0, len(heights)-1
	max := 0
	for i < j {
		max = maxVal(max, minVal(heights[i], heights[j])*(j-i))
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
