package main

import "fmt"

func main() {
	fmt.Println(wateringPlants([]int{2, 2, 3, 3}, 5))
}

func wateringPlants(plants []int, capacity int) int {
	res := 0

	now := capacity
	for i := 1; i <= len(plants); i++ {
		now -= plants[i-1]
		if i < len(plants) && now < plants[i] {
			res += i + i
			now = capacity
		}
		res++
	}

	return res
}
