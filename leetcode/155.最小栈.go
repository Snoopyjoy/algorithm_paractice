package leetcode.155

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

type ListItem struct {
	Value int
	Next *ListItem
	Prev *ListItem
}

type List struct {
	Head *ListItem
	Tail *ListItem
}

func (l *List)Push(v int){
	i := &ListItem{
		Value:v,
	}
	if l.Head == nil {
		l.Head = i
	}
	if l.Tail == nil {
		l.Tail = i
	} else {
		i.Prev = l.Tail
		l.Tail.Next = i
		l.Tail = i
	}
}

func (l *List)Pop() {
	if l.IsEmpty() {
		return
	}
	if l.Tail.Prev == nil {
		l.Tail = nil
		return
	}
	l.Tail, l.Tail.Prev.Next = l.Tail.Prev, nil
}

func (l *List)IsEmpty() bool {
	return l.Tail == nil
}

type MinStack struct {
	container *List
	helper *List
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		container: &List{},
		helper: &List{},
	}
}


func (this *MinStack) Push(x int)  {
	this.container.Push(x)
	lt := this.helper.Tail
	if lt == nil {
		this.helper.Push(x)
		return
	}
	if x <= lt.Value{
		this.helper.Push(x)
	}
}


func (this *MinStack) Pop()  {
	tail := this.container.Tail
	this.container.Pop()
	htail := this.helper.Tail
	if tail != nil && htail != nil && tail.Value == htail.Value{
		this.helper.Pop()
	}
}


func (this *MinStack) Top() int {
	if this.container.Tail == nil {
		return 0
	}
	return this.container.Tail.Value
}


func (this *MinStack) GetMin() int {
	if this.helper.Tail == nil {
		return 0
	}
	return this.helper.Tail.Value
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