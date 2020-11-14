/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	// 最大队列
	dequeue := list.New()

	var addToDequeue func(int)
	addToDequeue = func(idx int) {
		for dequeue.Back() != nil && nums[idx] > nums[dequeue.Back().Value.(int)] {
			dequeue.Remove(dequeue.Back())
		}
		dequeue.PushBack(idx)
	}
	for i:=0;i<k;i++{
		addToDequeue(i)
	}
	res := []int{nums[dequeue.Front().Value.(int)]}

	for i:=k; i<len(nums); i++{
		if dequeue.Front() != nil && dequeue.Front().Value.(int) <= i-k {
			dequeue.Remove(dequeue.Front())
		}
		addToDequeue(i)
		res = append(res, nums[dequeue.Front().Value.(int)])
	}
	return res
}
// @lc code=end
