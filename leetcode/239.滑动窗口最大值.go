package leetcode.239
/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 1 || k == 1{
		return nums
	}
	if k == 0 {
		return []int{}
	}
	left := make([]int, len(nums))
	left[0] = nums[0]
	right := make([]int, len(nums))
	right[len(nums)-1] = nums[len(nums)-1]
	for i:=1; i<len(nums); i++ {
		// 从左到右
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}
		// 从右到左
		j := len(nums)-i-1
		if (j + 1) %k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}
	ret := make([]int, 0, len(nums)-k)
	for i:=0; i < len(nums)-k+1; i++{
		ret = append(ret, max(left[i+k-1], right[i]))
	}
	return ret
}

func max(a, b int)int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) <= k {
		return []int{maxInArr(nums)}
	}
	// 窗口
	window := nums[:k]
	
	ret := []int{maxInArr(window)}

	for i := 1; i < len(nums) - k + 1; i++ {
		window = nums[i:i+k]
		//fmt.Println(window)
		ret = append(ret, maxInArr(window))
	}


	return ret
}

func maxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max{
			max = arr[i]
		}
	}
	return max
}


type ListItem struct {
	Prev  *ListItem
	Next  *ListItem
	Value int
}

type List struct {
	Head *ListItem
	Tail *ListItem
	Len  int
}

func (l *List) PopFront() *ListItem {
	if l.Len == 0 {
		return nil
	}
	ret := l.Head
	l.Len--
	if l.Len == 0 {
		l.Tail = nil
		l.Head = nil
		return ret
	}
	l.Head = l.Head.Next
	return ret
}

func (l *List) PopBack() *ListItem {
	if l.Len == 0 {
		return nil
	}
	ret := l.Tail
	l.Len--
	if l.Len == 0 {
		l.Tail = nil
		l.Head = nil
		return ret
	}
	l.Tail = l.Tail.Prev
	return ret
}

func (l *List) PushFront(x int) {
	l.Len++
	i := &ListItem{
		Value: x,
	}
	if l.Len == 1 {
		l.Head = i
		l.Tail = i
		return
	}
	l.Head, i, i.Next = i, l.Head.Prev, l.Head
}

func (l *List) PushBack(x int) {
	l.Len++
	i := &ListItem{
		Value: x,
	}
	if l.Len == 1 {
		l.Head = i
		l.Tail = i
		return
	}
	l.Tail, l.Tail.Next, i.Prev = i, i, l.Tail
}

type MaxQueue struct {
	container *List
	helper    *List
}

func (m *MaxQueue) In(x int) {
	m.container.PushBack(x)
	for m.helper.Len > 0 && m.helper.Tail.Value < x {
		m.helper.PopBack()
	}
	m.helper.PushBack(x)
}

func (m *MaxQueue) Out() {
	i := m.container.PopFront()
	if m.helper.Len > 0 && i.Value == m.helper.Head.Value {
		m.helper.PopFront()
	}
}

func (m *MaxQueue) Max() int {
	if m.helper.Len > 0 {
		return m.helper.Head.Value
	}
	return 0
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0, len(nums))
	window := &MaxQueue{
		container: &List{},
		helper:    &List{},
	}
	left, right := 0, 0
	for right < len(nums) {
		window.In(nums[right])
		right++
		if right-left == k {
			// 第一个窗口
			ret = append(ret, window.Max())

		}

		for right-left > k {
			left++
			window.Out()
			ret = append(ret, window.Max())
		}
	}
	return ret
}