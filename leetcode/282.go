package main

import "fmt"

//给定一个仅包含数字0-9的字符串 num 和一个目标值整数 target ，在 num 的数字之间添加 二元 运算符（不是一元
// +、-或*，返回所有能够得到目标值的表达式。
//
//
//示例 1:
//
//输入: num = "123", target = 6
//输出: ["1+2+3", "1*2*3"]

func main() {
	fmt.Println(addOperators("00", 0))
}

var Num string
var Target int
var lenn int
var res []string
var cal = []string{"+", "-", "*"}
var m map[string]bool

func addOperators(num string, target int) []string {
	Num = num
	Target = target
	lenn = len(num)
	m = make(map[string]bool)
	res = []string{}

	backtrace(0, 0, 0, 0, 0)

	return res
}

var now string
// pre -> all
// preNum -> one
func backtrace(pre, preNum, nowCal, nowNum, index int) {
	if index == lenn {
		if nowNum == Target {
			if !m[now] {
				m[now] = true
				res = append(res, now)
			}
		}

		return
	}

	nnum := 0
	nowStr := now
	if now != "" {
		now = now + cal[nowCal]
	}
	preOri := pre
	preNumOri := preNum
	nowNumOri := nowNum
	for i := index; i < lenn; i++ {
		if Num[index] == '0' && i != index {
			break
		}

		nnum = nnum*10 + int(Num[i]-'0')
		now = now + string(Num[i])
		switch nowCal {
		case 0:
			pre = nowNum
			preNum = nnum
			nowNum += nnum
		case 1:
			pre = nowNum
			preNum = - nnum
			nowNum -= nnum
		case 2:
			preNum = preNum * nnum
			nowNum = pre + preNum
		}

		if index+1 == lenn {
			backtrace(pre, preNum, 0, nowNum, i+1)
		} else {
			for j := 3; j > 0; j-- {
				backtrace(pre, preNum, j, nowNum, i+1)
			}
		}

		pre = preOri
		preNum = preNumOri
		nowNum = nowNumOri
	}
	now = nowStr
}
