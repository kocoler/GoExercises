package main

func totalMoney(n int) int {
	last := 0
	res := 0
	mon := 0
	now := 0
	for n > 0 {
		if mon == 0 {
			mon = 7
			last ++
			now = last
		}
		res += now
		now ++
		mon --
		n --
	}

	return res
}
