/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
import "container/list"

func maxSlidingWindow1(nums []int, k int) []int {
	// 最大队列
	dequeue := list.New()

	var addToDequeue func(int)
	addToDequeue = func(idx int) {
		for dequeue.Back() != nil && nums[idx] > nums[dequeue.Back().Value.(int)] {
			dequeue.Remove(dequeue.Back())
		}
		dequeue.PushBack(idx)
	}
	for i := 0; i < k; i++ {
		addToDequeue(i)
	}
	res := []int{nums[dequeue.Front().Value.(int)]}

	for i := k; i < len(nums); i++ {
		if dequeue.Front() != nil && dequeue.Front().Value.(int) <= i-k {
			dequeue.Remove(dequeue.Front())
		}
		addToDequeue(i)
		res = append(res, nums[dequeue.Front().Value.(int)])
	}
	return res
}

func maxSlidingWindow(num []int, size int) []int {
	if size > len(num) || len(num) == 0 {
		return []int{}
	}
	// 单调队列，单调递减
	dequeue := list.New()
	for i := 0; i < size; i++ {
		addToDequeue(dequeue, num, i)
	}
	if size == len(num) && dequeue.Front() != nil {
		return []int{num[dequeue.Front().Value.(int)]}
	}
	retindex := []int{dequeue.Front().Value.(int)}
	for i := size; i < len(num); i++ {
		if dequeue.Front() != nil && dequeue.Front().Value.(int) <= i-size {
			dequeue.Remove(dequeue.Front())
		}
		addToDequeue(dequeue, num, i)
		retindex = append(retindex, dequeue.Front().Value.(int))
	}

	ret := make([]int, 0, len(retindex))
	for _, v := range retindex {
		ret = append(ret, num[v])
	}
	return ret
}

func addToDequeue(dequeue *list.List, num []int, idx int) {
	for dequeue.Back() != nil && num[dequeue.Back().Value.(int)] < num[idx] {
		dequeue.Remove(dequeue.Back())
	}
	dequeue.PushBack(idx)
}

// @lc code=end
