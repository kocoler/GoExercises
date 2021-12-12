package main

// 这题完全不需要回溯诶
// 我就是个 sb

func circularPermutation(n int, start int) []int {
	nn = 1 << n
	N = n
	res = make([]int, nn)
	res[0] = start
	m = make([]bool, nn)
	m[start] = true
	find = false

	backtrace(1)

	return res
}

var res []int
var nn int
var N int
var m []bool
var find bool

func backtrace(bit int) {
	if bit == nn+1 {
		find = true
		return
	}
	last := res[bit-1]
	b := 1
	var val int
	for i := 0; i < N; i++ {
		t := b & last
		if t == 0 {
			val = last + (1 << i)
		} else {
			val = last - (1 << i)
		}
		if val < nn && !m[val] {
			res[bit] = val
			m[val] = true
			backtrace(bit + 1)
			if find {
				return
			}
		}
		b <<= 1
	}
}
