// 还是 众数 问题，也就是 摩尔投票法

func majorityElement(nums []int) int {
    candidate := 0
    count := 0

    for _, v := range nums {
        if candidate == v {
            count ++
        } else {
            count --
        }

        if count <= 0 {
            candidate = v
            count = 1
        }
    }

    return candidate
}