package main

import"fmt"

func main(){
	a :=make([]int,100)
	a ={1,2,3}
	b :=[]int {4,5,6}
	a =append(a[0:2],b[0:]...)
	fmt.Println(a)
}
