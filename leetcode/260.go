// 分组异或位运算
// 第一组，第二组
// 选取其中的一位，另外不同的必定是 0

func singleNumber(nums []int) []int {
    var p int
    for _, v := range nums {
        p = p ^ v
    }

    var b int
    for i := 0; i < 32; i ++ {
        if (p >> i) & 1 == 1 {
            b = 1 << i
            break
        }
    }

    res := make([]int, 2)

    for _, v := range nums {
        if b & v != 0 {
            res[0] = res[0] ^ v
        } else {
            res[1] = res[1] ^ v
        }
    }

    return res
}
