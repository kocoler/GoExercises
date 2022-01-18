package main

func superPow(a int, b []int) int {
	mod := 1337
	res := 1
	quickPower := func(x, n int) int {
		ret := 1
		for n > 0 {
			if n&1 == 1 {
				ret = ret * x % mod
			}
			n >>= 1
			x = x * x % mod
		}

		return ret
	}

	for i := len(b) - 1; i >= 0; i-- {
		res = res * quickPower(a, b[i]) % mod
		a = quickPower(a, 10)
	}

	return res
}
