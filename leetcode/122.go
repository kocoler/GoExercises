package main


// 贪心
func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    sum := 0 
    tmp := 0
    for i := 0 ; i < len(prices)-1 ; i++ {
        if prices[i] < prices[i+1] {
            if tmp == 0 {
                sum -= prices[i]
                tmp = 1
                continue
            }
        }   
        if prices[i] > prices[i+1] {
            if tmp == 1 {
                for j := i ; j < len(prices)-1 ; j++ {
                if prices[j] > prices[j+1] {
                    i = j
                    sum += prices[j]
                    tmp = 0
                    break
                }
                }
            }  
        }
    }
    if tmp ==1 {
        sum += prices[len(prices)-1]
    }
    return sum
}

// 动规
var dp[30001][2]int
func maxProfit(prices []int) int {
    // 0 num 0
    // 1 num 1
    dp = [30001][2]int{}

    dp[0][0] = 0
    dp[0][1] = -prices[0]

    lenp := len(prices)

    for i := 1; i < lenp; i ++ {
        index := i
        // out
        dp[i][0] = max(dp[i-1][1]+prices[index], dp[i-1][0])
        // in
        dp[i][1] = max(dp[i-1][0]-prices[index], dp[i-1][1])
    }

    return max(dp[lenp-1][0], dp[lenp-1][1])
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

