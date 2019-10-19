
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    sum :=len(nums1)+len(nums2)
		var a =make([]int,10000)
		var num float64
		b:=0;c:=0
		for i :=0;i<sum/2+1;i++ {
		if len(nums1)==0 {
		a=nums2
		break
	}
		if len(nums2)==0 {
		a=nums1
		break
	}
		if(nums1[b]<=nums2[c]){
		a[i]=nums1[b]
		b++
		if b>=len(nums1) {
		a=append(a[0:i+1],nums2[c:]...)
		break
	}
	}else {
		a[i]=nums2[c]
		c++
		if c>=len(nums2) {
			a=append(a[0:i+1],nums1[b:]...)
			break
		}
	}
	}
		if(sum%2!=0){
		num=float64(a[sum/2])
	}else{
		num=float64(a[sum/2-1]+a[sum/2])
		num=num/2
	}
    return num
}


 

