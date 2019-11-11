package main
import "fmt"
/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	Head *LinkedList
	Tail *LinkedList
	Length int
	Size int
}

type LinkedList struct {
	Val int
	Prev *LinkedList
	Next *LinkedList
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	res := MyCircularDeque{
		Head: nil,
		Tail: nil,
		Length: 0,
		Size: k}
	return res
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
		return false
	}
	newNode := &LinkedList{
		Val: value,
		Prev: nil,
		Next: nil }
	if this.IsEmpty() {
		this.Tail = newNode
		this.Head = newNode
	}else{
		newNode.Next = this.Head
        this.Head.Prev = newNode
        this.Head = newNode
	}
    this.Length++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
		return false
	}
    newNode := &LinkedList{
		Val: value,
		Prev: nil,
		Next: nil}
    if( this.IsEmpty() ){
        this.Head = newNode
        this.Tail = newNode
    }else{
        newNode.Prev = this.Tail
        this.Tail.Next = newNode
        this.Tail = newNode
    }
    this.Length++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
		return false
	}
    this.Head = this.Head.Next
    this.Length--
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty(){
		return false
	}
    this.Tail = this.Tail.Prev
    this.Length--
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
		return -1
	}
    return this.Head.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
		return -1
	}
    return this.Tail.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.Length == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.Length == this.Size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

func main(){
	obj := Constructor(2)
	param_1 := obj.InsertFront(1)
	fmt.Println(param_1)
	param_2 := obj.InsertLast(2)
	fmt.Println(param_2)
	param_3 := obj.DeleteFront()
	fmt.Println(param_3)
	param_4 := obj.DeleteLast()
	fmt.Println(param_4)
	param_5 := obj.GetFront()
	fmt.Println(param_5)
	param_6 := obj.GetRear()
	fmt.Println(param_6)
	param_7 := obj.IsEmpty()
	fmt.Println(param_7)
	param_8 := obj.IsFull()
	fmt.Println(param_8)
}


