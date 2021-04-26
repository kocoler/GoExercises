package main

func jump(nums []int) int {
    if len(nums)==1 {
        return 0
    }
    sum := 0
	for i := 0; i<len(nums); i++{
		max := 0
		n := i 
		if nums[i]+i>=len(nums)-1{
			sum++
			break
		}
		for j := 1;j<=nums[i]; j++{
			if (nums[i+j]+j)>=max {
				n = j+i
				max = nums[n]+j
			}
		}
		i = n-1
		sum++
	}
    return sum
}


 

