package main

import "fmt"

func main() {
	var res []int
	res = append(res, 0)
	for i := 0; i < 3; i ++ {
		res = append(res, 0)
	}
	//for i := 0; i < 1; i ++ {
	//	res = append(res, 1)
	//}
	for i := 0; i < 9; i ++ {
		res = append(res, 1)
	}
	//for i := 0; i < 1; i ++ {
	//	res = append(res, 1)
	//}
	for i := 0; i < 3; i ++ {
		res = append(res, 0)
	}
	res = append(res, 0)
	fmt.Println(len(res))
	fmt.Print("[")
	for i := 0; i < len(res); i ++ {
		fmt.Print(res[i], ",")
	}
	fmt.Print("]")
}
