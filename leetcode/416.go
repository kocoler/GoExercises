package main

func canPartition(nums []int) bool {
	lenn := len(nums)

	sum := 0
	for _, v := range nums {
		sum += v
	}

	target := sum / 2
	if target * 2 != sum {
		return false
	}

	dp := make([]bool, target+1)

	dp[0] = true
	for i := 0; i < lenn; i ++ {
		value := nums[i]
		for j := target; j >= value; j -- {
			dp[j] = dp[j-value] || dp[j]
		}
	}

	return dp[target]
}
