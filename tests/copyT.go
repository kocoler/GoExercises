package main

import "fmt"

func main() {
	str1 := []string{"abc"}
	str2 := []string{"def"}
	i := copy(str1, str2)
	fmt.Println(i, str1, str2)
	str1 = str2
	fmt.Printf("%p, %p", &str1, &str2)
}
