package main

// rand7 -> rand10
// 思路：生成 >=10 的规模，取其中的前xxx个
// 公式：(randx - 1) * x + randx
// 优化：反馈继续缩小抛弃范围（拒绝采样）

func rand10() int {
	for {
		a := (rand7() - 1) * 7 + rand7()
		if (a <= 40) {
			return a % 10 + 1
		}
		a = (a - 40 - 1) * 7 + rand7()
		if (a <= 60) {
			return a % 10 + 1
		}
		a = (a - 60 - 1) * 7 + rand7()
		if (a <= 20) {
			return 1 + a % 10
		}
	}

	return 0
}

