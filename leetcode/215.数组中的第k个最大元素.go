/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return findKth(nums, 0, len(nums)-1, len(nums)-k)
}

func findKth(nums []int, i, j, k int) int {
	p := partition(nums, i, j)
	if p == k {
		return nums[p]
	} else if p < k {
		return findKth(nums, p+1, j, k)
	} else {
		return findKth(nums, i, p-1, k)
	}
}

func partition(arr []int, i, j int) int {
	mark := i
	for k := i + 1; k <= j; k++ {
		if arr[k] < arr[i] {
			mark++
			arr[mark], arr[k] = arr[k], arr[mark]
		}
	}
	arr[i], arr[mark] = arr[mark], arr[i]
	return mark
}

// @lc code=end

