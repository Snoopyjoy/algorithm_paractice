import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	i, j, k := 0, 1, len(nums)-1
	ret := nums[0] + nums[1] + nums[2]
	for i < len(nums)-2 {
		curT := target - nums[i]
		for j < k {
			v := nums[j] + nums[k]
			tt := nums[i] + v

			if abs(target-tt) < abs(target-ret) {
				ret = tt
			}
			if v > curT {
				k--
			} else if v < curT {
				j++
			} else {
				return target
			}
		}

		i++
		j = i + 1
		k = len(nums) - 1
	}

	return ret
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

