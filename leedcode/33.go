package main

func search(nums []int, target int) int {
    if len(nums)==0{
        return -1
    }
    mid := len(nums)/2
    
    left := 0;right := len(nums)-1
    for right>left {
        if nums[mid]==target {
        return mid
        }
        if nums[mid]>nums[left]{
            if target<=nums[left]{
                left = mid
                mid = (left+right)/2
            }
            if target<=nums[mid]{
                
            }
        }
    }
    return -1
}


 

