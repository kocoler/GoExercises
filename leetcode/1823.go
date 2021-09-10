// 约瑟夫环
// 倒序计算索引 (res + k) % i

func findTheWinner(n int, k int) int {
    res := 0
    for i := 2; i <= n; i ++ {
        res = (res + k) % i
    }
    return res + 1
}