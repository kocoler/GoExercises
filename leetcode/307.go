package main

// 树状数组
// 简单求区间和，用「前缀和」
// 多次将某个区间变成同一个数，用「线段树」
// 其他情况，用「树状数组」

type treeArray struct {
	tree []int
	nums []int
	n    int
}

func (t *treeArray) lowbit(x int) int {
	return x & (-x)
}

func (t *treeArray) query(x int) int {
	var res int

	for x > 0 {
		res += t.tree[x]
		x -= t.lowbit(x)
	}

	return res
}

func (t *treeArray) add(x, u int) {
	i := x

	for i <= t.n {
		t.tree[i] += u
		i += t.lowbit(i)
	}
}

func (t *treeArray) Build(nums []int) {
	t.tree = make([]int, len(nums)+1)
	t.nums = nums
	t.n = len(nums)

	for i := 0; i < len(nums); i++ {
		t.add(i+1, nums[i])
	}
}

func (t *treeArray) Update(x, val int) {
	t.add(x+1, val-t.nums[x])
	t.nums[x] = val
}

func (t *treeArray) SumRange(l, r int) int {
	return t.query(r+1) - t.query(l)
}

type NumArray struct {
	treeArray treeArray
}

func Constructor(nums []int) NumArray {
	var na NumArray

	na.treeArray.Build(nums)

	return na
}

func (this *NumArray) Update(index int, val int) {
	this.treeArray.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.treeArray.SumRange(left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
