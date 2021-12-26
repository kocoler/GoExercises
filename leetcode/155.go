package main

type MinStack struct {
	ori []int
	// 最小栈
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.ori = append(this.ori, val)
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= val {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	x := this.ori[len(this.ori)-1]
	this.ori = this.ori[:len(this.ori)-1]
	if this.minStack[len(this.minStack)-1] == x {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.ori[len(this.ori)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
