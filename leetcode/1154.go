package main

import "strconv"

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])

	//能被400整除的，或者不能被100整除而能被4整除的年就是闰年，换一句话说，非世纪年份中能被4整除的，和世纪年份中能被400整除的是闰年。按照这个定义，公元2000年是闰年，而公元1900年是平年。
	r := year%400 == 0 || (year%100 != 0 && year%4 == 0)

	res := 0
	for i := 1; i <= month; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			res += 31
		case 2:
			if r {
				res += 28
			} else {
				res += 29
			}
		default:
			res += 30
		}
	}

	return res
}
