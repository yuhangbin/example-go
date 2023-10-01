package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	// 1. 0 0 0
	// 2. - - +
	// 3. + + -
	// 4. + - 0
	n := len(nums)
	var result [][]int
	if n < 3 {
		return result
	}
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}

			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result

}
