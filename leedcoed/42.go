
func trap(height []int) int {
    var num,leftm,rightm int
    left := 0;right := len(height)-1
    for left<right {        
        if height[left]<height[right] {
            if leftm<height[left] {
            leftm = height[left]
            }
            num+=leftm-height[left]
            left++
        }else{
            if rightm<height[right] {
            rightm = height[right]
            }
            num+=rightm-height[right]
            right--
        }
    }
    return num
}


 

