package main

import "fmt"

func mergesort(arr []int) []int{
	if(len(arr)<2){
		return arr
	}
	num:=len(arr)/2
	left:=mergesort(arr[0:num])
	right:=mergesort(arr[num:])
	fin:=merge(left,right)
	return fin
}

func merge(left,right []int) (result []int){
	l,r:=0,0
	for l<len(left)&&r<len(right) {
		    if (left[l]<right[r]) {
                result=append(result,left[l])
                l++
        } else {
                result=append(result,right[r])
                r++
        }
}

	result=append(result,left[l:]...)
	result=append(result,right[r:]...)
	return result
}

func main(){a:=[]int{4,2,5,7,1,2}
	a=mergesort(a)
	fmt.Println(a)
}
