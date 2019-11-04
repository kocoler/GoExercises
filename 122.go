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
