package main

import "fmt"

type cub struct {
	l int
	d int
	h int
}

func (cub1 *cub) v()  {
	fmt.Printf("长方体的体积为：%d",cub1.l*cub1.d*cub1.h)
}

func (cub1 *cub) s()  {
	fmt.Printf("长方体的表面积为：%d",(cub1.l*cub1.h)*2+(cub1.l*cub1.d)*2+(cub1.d*cub1.h)*2)
}

func main()  {
	var cubx cub
	fmt.Printf("输入长：")
	fmt.Scanf("%d",&cubx.l)
	fmt.Printf("输入宽：")
	fmt.Scanf("%d",&cubx.d)
	fmt.Printf("输入高：")
	fmt.Scanf("%d",&cubx.h)
	cubx.v()
	cubx.s()
}
