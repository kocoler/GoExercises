package main

import (
	"fmt"
)

// 罗马数字转阿拉伯数字
func RomanToInt(romanNum string) int {
	num := 0
	for i := len(romanNum) - 1; i >= 0; i-- {
		switch {
		case romanNum[i] == 'I':
			num++
		case romanNum[i] == 'V':
			if i >= 1 && romanNum[i-1] == 'I' {
				num += 4
				i--
			} else {
				num += 5
			}
		case romanNum[i] == 'X':
			if i >= 1 && romanNum[i-1] == 'I' {
				num += 9
				i--
			} else {
				num += 10
			}
		case romanNum[i] == 'L':
			if i >= 1 && romanNum[i-1] == 'X' {
				num += 40
				i--
			} else {
				num += 50
			}
		case romanNum[i] == 'C':
			if i >= 1 && romanNum[i-1] == 'X' {
				num += 90
				i--
			} else {
				num += 100
			}
		case romanNum[i] == 'D':
			if i >= 1 && romanNum[i-1] == 'C' {
				num += 400
				i--
			} else {
				num += 500
			}
		case romanNum[i] == 'M':
			if i >= 1 && romanNum[i-1] == 'C' {
				num += 900
				i--
			} else {
				num += 1000
			}
			//default: continue
		}
	}
	return num
}

// 阿拉伯数字转罗马数字
func IntToRoman(num int) []string {
	c := make(map[int]string)
	//j:=10
	var roma []string
	c[1] = "I"
	c[2] = "II"
	c[3] = "III"
	c[4] = "IV"
	c[5] = "V"
	c[6] = "VI"
	c[7] = "VII"
	c[8] = "VIII"
	c[10] = "X"
	c[9] = "IX"
	c[20] = "XX"
	c[30] = "XXX"
	c[40] = "XL"
	c[50] = "L"
	c[60] = "LX"
	c[70] = "LXX"
	c[80] = "LXXX"
	c[90] = "XC"
	c[100] = "C"
	c[200] = "CC"
	c[300] = "CCC"
	c[600] = "DC"
	c[700] = "DCC"
	c[800] = "DCCC"
	c[1000] = "M"
	c[900] = "CM"
	c[500] = "D"
	c[400] = "CD"
	for i := 1; i <= 1000; i = i * 10 {
		roma = append(roma, c[num%10*i])
		num = num / 10
		if num == 0 {
			break
		}
	}
	/*if num>0{
		//fmt.Printf("!!")
		for i:=num%10*j;i>0;i--{
			//fmt.Printf("!!")
			roma=append(roma,c[1000])
		}
		j=j*10
		num=num/10
	}*/
	return roma
}

func main() {
	var a []string
	var b, n, e int
	var d string
	fmt.Printf("1：罗马数字，2：阿拉伯数字")
	fmt.Scanf("%d", &n)
	if n == 2 {
		fmt.Scanf("%d", &b)
		a = IntToRoman(b)
		for i := len(a) - 1; i >= 0; i-- {
			fmt.Printf("%s", a[i])
		}
	} else {
		fmt.Scanf("%s", &d)
		e = RomanToInt(d)
		fmt.Printf("%d", e)
	}
}
