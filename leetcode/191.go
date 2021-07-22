package main

func hammingWeight(num uint32) int {
	var res uint32

	for num != 0 {
		res += num & 1
		num >>= 1
	}

	return int(res)
}

// n & (n - 1) => 减少最后的 一

func hammingWeight(num uint32) int {
	var res int

	for ; num != 0; num &= num - 1{
		res ++
	}

	return res
}
