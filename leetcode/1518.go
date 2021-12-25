package main

func numWaterBottles(numBottles int, numExchange int) int {
	nowEmpytNum := numBottles
	res := numBottles
	for nowEmpytNum >= numExchange {
		ex := nowEmpytNum / numExchange
		nowEmpytNum %= numExchange
		nowEmpytNum += ex
		res += ex
	}
	return res
}
