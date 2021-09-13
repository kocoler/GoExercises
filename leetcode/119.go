// 倒序滚动数组

func getRow(rowIndex int) []int {
    res := make([]int, rowIndex + 1)

    res[0] = 1
    for i := 1; i <= rowIndex; i ++ {
        for j := i; j >= 1; j -- {
            res[j] = res[j] + res[j - 1]
        }
    }

    return res
}