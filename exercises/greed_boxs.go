package main

import "fmt"

type detail struct{
	w float32
	s float32
	v float32
}

func main()  {
	var a[100] detail
	var tmp detail
	fmt.Printf("輸入")
	var k,n int
	var m float32
	fmt.Scanf("%d",&k)
	for k>0 {
		var sum float32;sum = 0//float32;sum = 0.0
		fmt.Scanf("%f",&m)
		fmt.Scanf("%d",&n)
		for i:=0;i<n;i++{
			fmt.Scanf("%f",&a[i].w)
			fmt.Scanf("%f",&a[i].s)
			a[i].v=a[i].s/a[i].w
		}
		for i:=0;i<n;i++{
			for j:=i+1;j<n;j++ {
				if a[i].v<a[j].v {
					tmp=a[i]
					a[i]=a[j]
					a[j]=tmp
				}
			}
		}
		for i:=0;i<n;i++  {
			if m>=a[i].w {
				sum=a[i].s+sum
				m=m-a[i].w
			}else {
				sum=sum+a[i].v*m
				break
			}
		}
		fmt.Printf("%.2f",sum)
		k--
	}
}

