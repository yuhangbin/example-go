package leetcode

func MaxArea(height []int) int {
	head, tail := 0, len(height)-1

	max := 0

	for head < tail {
		max = Max(max, (tail-head)*Min(height[tail], height[head]))
		if height[tail] < height[head] {
			tail = tail - 1
		} else {
			head = head + 1
		}
	}
	return max
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
