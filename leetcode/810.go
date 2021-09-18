package main

// 博弈论
// 假定奇偶分析法(因为是交替状态)
// 先确定一个 奇/偶 状态，再推出他 输/赢 的可能性
// 那么就能确定这个 奇/偶 状态是否会 输/赢

func xorGame(nums []int) bool {
	s := 0
	for _, v := range nums {
		s ^= v
	}

	return s == 0 || len(nums) % 2 == 0
}
