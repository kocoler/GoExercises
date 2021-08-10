package lcof

// 双栈
// 一个只出
// 一个只入
// 只入的再放到只出的里，这样就形成了正序的队列

type CQueue struct {
	stackIntsF []int
	stackIntsT []int
}

func (c *CQueue)popFront() (res int) {
	if len(c.stackIntsF) + len(c.stackIntsT) < 1 {
		return -1
	}

	if len(c.stackIntsT) == 0 {
		lenf := len(c.stackIntsF)
		for i := 1; i <= lenf; i ++ {
			c.stackIntsT = append(c.stackIntsT, c.stackIntsF[lenf - i])
		}
		c.stackIntsF = []int{}
	}


	res = c.stackIntsT[len(c.stackIntsT) -1]
	c.stackIntsT = c.stackIntsT[:len(c.stackIntsT) - 1]
	return
}

func (c *CQueue)pushBack(i int) {
	c.stackIntsF = append(c.stackIntsF, i)
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.pushBack(value)
}


func (this *CQueue) DeleteHead() int {
	return this.popFront()
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */