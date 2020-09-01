package leetcode_15

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	ret := [][]int{}
	var a, b, c int
	var target int
	var j, k int
	for i := 0; i < len(nums)-2; i++ {
		a = nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}
		target = -a
		j = i + 1
		k = len(nums) - 1
		for j < k {
			b = nums[j]
			c = nums[k]
			if c < 0 {
				break
			}
			if b+c == target {
				ret = append(ret, []int{a, b, c})

				for j < k && nums[j+1] == b {
					j++
				}
				for j < k && nums[k-1] == c {
					k--
				}
				j++
				k--
				continue
			}
			if b+c < target {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}

// @lc code=end
