package lcof

type MaxQueue struct {
	queue []int
	len int
	maxQ []int
}


func Constructor() MaxQueue {
	return MaxQueue{}
}


func (this *MaxQueue) Max_value() int {
	if this.len == 0 {
		return -1
	}

	return this.maxQ[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.len ++
	this.queue = append(this.queue, value)
	lenMQ := len(this.maxQ)
	for i := lenMQ - 1; i >= 0; i -- {
		if this.maxQ[i] >= value {
			break
		}
		this.maxQ = this.maxQ[:lenMQ-1]
		lenMQ --
	}
	this.maxQ = append(this.maxQ, value)
}


func (this *MaxQueue) Pop_front() int {
	if this.len == 0 {
		return -1
	}

	this.len --
	res := this.queue[0]
	this.queue = this.queue[1:]
	//for i, v := range this.maxQ {
	//	if v == res {
	//		this.maxQ = append(this.maxQ[:i], this.maxQ[i+1:]...)
	//	}
	//}
	// 这样更好！
	// 仔细理解一下单调队列
	// 对于 res 入队时来说，所有比他小的，都出队了，比他大的，一定比他先入队
	// 所以到 res 出队时，所有比他大的，一定已经出队，如果他次时不是最大的，那么他也已经出队
	if this.maxQ[0] == res {
		this.maxQ = this.maxQ[1:]
	}

	return res
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
