package main

import "fmt"

func main() {
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for tickets[k] != 0 {
		// k
		for key := range tickets {
			if tickets[key] != 0 {
				res++
				tickets[key]--
			}
			if k == key && tickets[k] == 0 {
				break
			}
		}
	}

	return res
}
