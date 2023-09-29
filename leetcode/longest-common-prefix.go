package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		current := strs[0][i]
		for _, value := range strs {
			if len(value) <= i {
				return result.String()
			}
			if value[i] != current {
				return result.String()
			}
		}
		result.WriteString(string(current))
	}
	return result.String()
}
