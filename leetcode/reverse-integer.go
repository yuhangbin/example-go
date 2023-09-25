package leetcode

func Reverse(x int) int {
	result := 0
	max := (1 << 31) - 1
	min := -(max + 1)

	for x != 0 {
		tail := x % 10
		newResult := result*10 + tail
		if newResult > max || newResult < min {
			return 0
		}
		result = newResult
		x = x / 10
	}

	return result
}
