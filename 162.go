func findPeakElement(nums []int) int {
    //a := 0
    /*if len(nums)==2 {
        if nums[0] > nums[1] {
            return nums[0]
        }else {
            return nums[1]
        }
    }*/
    left := 0 ; right := len(nums)-1
    //mid := left + (right-left)/2
    for left < right {
        /*if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
        return1 mid
        }*/ 
        mid := (right + left)/2
        if nums[mid] < nums[mid+1] {
            left = mid + 1
        }else {
            right = mid
        }
    }
    return left
}
