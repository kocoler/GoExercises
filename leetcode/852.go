package main

// 山脉数组的峰顶索引
// 注意 left <= right
func peakIndexInMountainArray(arr []int) int {
    // mid - 1 < mid
    // mid + 1 > mid
    left := 0
    right := len(arr) - 1
    mid := (left + right) / 2
    for left <= right {
        mid = (left + right) / 2
        if arr[mid] > arr[mid+1] {
            if arr[mid] > arr[mid-1] {
                return mid
            }
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return mid
}

