// 用 int32 表示最终的数
// 每位统计所有的数字的 i 位总和，除3之后剩余的就说明多出来的这个数，是 res 的
// 再复原回去

func singleNumber(nums []int) int {
    var res int32

    for i := 0; i < 32; i ++ {
        var temp int32
        for _, v := range nums {
            temp =  temp + int32(v) >> i & 1
        }
        if temp % 3 != 0 {
            res |= 1 << i
        }
    }

    return int(res)
}