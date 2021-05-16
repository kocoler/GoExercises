package main

type ints []int

type FindSumPairs struct {
	Nums1 []int
	Nums2 []int
	M1 map[int][]int
	M2 map[int][]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m1 := make(map[int][]int)
	m2 := make(map[int][]int)
	for _, v := range nums1 {
		m1[v] = append(m1[v], v)
	}
	for _, v := range nums2 {
		m2[v] = append(m2[v], v)
	}
	return FindSumPairs{Nums1: nums1, Nums2: nums2, M1: m1, M2: m2}
}


func (this *FindSumPairs) Add(index int, val int)  {
	p := this.Nums2[index]
	vs := this.M2[this.Nums2[index]]
	this.Nums2[index] += val
	for k := range vs {
		if vs[k] == index {
			this.M2[p] = append(vs[:k], vs[k+1:] ...)
			break
		}
	}
	this.M2[this.Nums2[index]] = append(this.M2[this.Nums2[index]], index)
}


func (this *FindSumPairs) Count(tot int) int {
	res := 0
	for k, v := range this.M1 {
		if k > tot {
			return res
		}
		lenv := len(v)
		lenv2 := len(this.M2[tot-k])
		res += lenv*lenv2
	}

	return res
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
