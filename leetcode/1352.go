package main

// 要注意 0
// 乘到 0 了就得清空哦~

type ProductOfNumbers struct {
	preMulti []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		preMulti: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.preMulti = []int{1}
		return
	}
	this.preMulti = append(this.preMulti, this.preMulti[len(this.preMulti)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.preMulti) <= k {
		return 0
	}

	return this.preMulti[len(this.preMulti)-1] / this.preMulti[len(this.preMulti)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
