package leetcode

func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, n, 0, 0)
	return res
}

func backtrack(res *[]string, s string, n, open, close int) {
	if open == n && close == n {
		*res = append(*res, s)
		return
	}

	if open < n {
		backtrack(res, s+"(", n, open+1, close)
	}

	if close < open {
		backtrack(res, s+")", n, open, close+1)
	}
}
