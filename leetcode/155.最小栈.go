package leetcode.155

import (
	"fmt"
	"container/list"
)

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

type MinStack struct {
	container *list.List
	helper *list.List
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		container: list.New(),
		helper: list.New(),
	}
}

func (this *MinStack) Push(x int)  {
	this.container.PushBack(x)
	if this.helper.Len() == 0 {
		this.helper.PushBack(x)
		return
	}
	bv, _ := this.helper.Back().Value.(int)
	if x > bv {
		return
	}
	this.helper.PushBack(x)

}


func (this *MinStack) Pop()  {
	if this.container.Len() == 0 {
		return
	}
	bv, _ := this.container.Back().Value.(int)
	if this.helper.Len() == 0 {
		return
	}
	hbv, _ := this.helper.Back().Value.(int)
	this.container.Remove(this.container.Back())
	if bv == hbv {
		this.helper.Remove(this.helper.Back())
	}

}


func (this *MinStack) Top() int {
	if this.container.Len() == 0 {
		return 0
	}
	bv, _ := this.container.Back().Value.(int)
	return bv
}


func (this *MinStack) GetMin() int {
	if this.helper.Len() == 0 {
		return 0
	}
	bv, _ := this.helper.Back().Value.(int)
	return bv

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end