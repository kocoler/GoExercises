package main

// 差分

func corpFlightBookings(bookings [][]int, n int) []int {
	lenb := len(bookings)
	diff := make([]int, n)

	for i := 0; i < lenb; i ++ {
		// start: bookings[i][0]
		// end: bookings[i][1]
		start := bookings[i][0] - 1
		end := bookings[i][1] - 1
		value := bookings[i][2]
		diff[start] += value
		if end + 1 < n {
			diff[end + 1] -= value
		}
	}

	pre := diff[0]
	for i := 1; i < n; i ++ {
		diff[i] = diff[i] + pre
		pre = diff[i]
	}

	return diff
}
