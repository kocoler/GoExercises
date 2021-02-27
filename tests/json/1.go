package main

import (
	"encoding/json"
	"fmt"
)

type a struct {
	a int
}

func main() {
	var b []a
	bytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes), len(bytes), len(b))
}
