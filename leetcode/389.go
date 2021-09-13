// 也可以转化成位运算
// 当计算 xx 多一位 / 单独 / 重复 等等，可以考虑位运算

func findTheDifference(s string, t string) byte {
    l := s + t
    lenl := len(l)

    var res byte
    for i := 0; i < lenl; i ++ {
        res = res ^ l[i]
    }

    return res
}