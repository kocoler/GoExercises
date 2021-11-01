package main

import "fmt"

func main() {
	fmt.Println(exchangeBits(156312512))
}

func exchangeBits(num int) int {
	bit := 0b01010101010101010101010101010101

	// (num & bit) << 1 保留了num中的奇数位并与偶数位交换
	// (num & bit << 1) >> 1 保留了num中的偶数位并与奇数位交换

	return (num&bit)<<1 | (num&(bit<<1))>>1
}
