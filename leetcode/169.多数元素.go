package leetcode_169

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	counter := make(map[int]int, len(nums))
	pos := make([]int, len(nums))

	for _, v := range nums {
		counter[v]++
		pos[counter[v]-1] = v
	}
	return pos[len(nums)/2]
}
func majorityElement1(nums []int) int {
	k := len(nums) / 2
	return findKth(nums, 0, len(nums)-1, k)
}

func findKth(nums []int, start, end, k int) int {
	n := partition(nums, start, end)
	if n == k {
		return nums[n]
	}
	if n > k {
		return findKth(nums, start, n-1, k)
	}
	return findKth(nums, n+1, end, k)
}

func partition(nums []int, start, end int) int {
	mark := start
	for i := start + 1; i <= end; i++ {
		if nums[i] < nums[start] {
			mark++
			nums[i], nums[mark] = nums[mark], nums[i]
		}
	}
	nums[start], nums[mark] = nums[mark], nums[start]
	return mark
}

// @lc code=end
