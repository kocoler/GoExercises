package main

import "fmt"

func main() {
	maps := make(map[int]int)
	maps[1] = 1
	fmt.Println(len(maps))
	maps[1] --
	fmt.Println(len(maps))
}
