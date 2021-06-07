package main

import "fmt"

func main() {
	ch := make(chan error, 1)

	close(ch)

	fmt.Println(<- ch)
}
