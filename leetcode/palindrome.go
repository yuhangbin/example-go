package leetcode

func IsPalindrome(x int) bool {
	y := x
	res := 0
	for y != 0 {
		res = res * 10 + y%10
		y/=10
	}
	return x == res
}
