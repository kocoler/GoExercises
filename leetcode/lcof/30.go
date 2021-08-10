package lcof

type MinStack struct {
	// 栈
	vals []int
	// 最小栈
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.vals = append(this.vals, x)

	lenm := len(this.minStack)
	if lenm == 0 || x <= this.minStack[lenm - 1] {
		this.minStack = append(this.minStack, x)
	}
}


func (this *MinStack) Pop()  {
	x := this.vals[len(this.vals) - 1]
	this.vals = this.vals[:len(this.vals) - 1]
	lenm := len(this.minStack)
	// i := lenm - 1
	// for ; i >= 0; i -- {
	//     if this.minStack[i] == x {
	//         this.minStack = append(this.minStack[:i], this.minStack[i+1:] ...)
	//     }
	// }
	if this.minStack[lenm - 1] == x {
		this.minStack = this.minStack[:lenm - 1]
	}
}


func (this *MinStack) Top() int {
	return this.vals[len(this.vals) - 1]
}


func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */