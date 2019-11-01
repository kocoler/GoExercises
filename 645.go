func findErrorNums(nums []int) []int {
    var a []int
    b := make(map[int]int)
    for i := 0 ;i < len(nums) ;i++ {
        if b[nums[i]]  == 1 {
            a = append(a,nums[i])
        }
         if b[nums[i]] != 1 {
             b[nums[i]] = 1
        }
    }
    for j := 1 ;j < len(nums)+1 ;j++ {
        if b[j] != 1 {
            a = append(a,j)
            return a
        }
    }
    return a
}
