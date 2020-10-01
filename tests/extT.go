package main

import (
	"fmt"
	"path"
)

func main() {
	fileName := "aajpg"
	fmt.Println(path.Ext(fileName))
}
