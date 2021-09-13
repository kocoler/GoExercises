package main

// 好久没有双百了，还是道困难题 哭泣
// 就是逆康托展开

var FAC = []int{1,1,2,6,24,120,720,5040,40320,362800}

func decantor(num int, n int) []byte {
	a := make([]byte, num)
	for i := 1; i <= num; i ++ {
		a[i-1] = byte('0' + i)
	}
	lena := len(a)
	res := make([]byte, num)

	for i := lena - 1; i >= 0; i -- {
		index := n / FAC[i]
		n = n % FAC[i]
		res[lena - 1 - i] = a[index]
		a = append(a[:index], a[index+1:] ...)
	}

	return res
}

func getPermutation(n int, k int) string {
	return string(decantor(n , k - 1))
}
