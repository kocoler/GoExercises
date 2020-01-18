package main

import (
	"fmt"
	"regexp"
)

type result struct{
	final  string
	daily string
	total string
}

func main() {
	str := `
		`
	reg1 := regexp.MustCompile("<td valign=\"middle\">(.*?)</td>")
	result1 := reg1.FindAllStringSubmatch(str,-1)
	//fmt.Println(result1[0][1])
	var i = 0
	fmt.Println(len(result1))
	type ResultDetails struct {
		ClassName  string
		DailyScore string
		DailyScoreRate string
		FinalScore string
		FinalScoreRate string
		TotalScore string
	}
	var result = ResultDetails{
		ClassName:      "1",
		DailyScore:     "2",
		DailyScoreRate: "3",
		FinalScore:     "4",
		FinalScoreRate: "5",
		TotalScore:     "6",
	}
	for _, text := range result1 {
		if i == 1 {
		   result.DailyScoreRate = ChangeString(text[1], 1, len(text[1])-5)
		   i = 2
		   continue
		}
		if i == 2 {
			result.DailyScore = ChangeString(text[1], 1, len(text[1])-5)
			i = 0
		}
		if i == 3 {
			//fmt.Println(text[1],len(text[1]),ChangeString(text[1], 1, len(text[1])-6) )
			result.FinalScoreRate = ChangeString(text[1], 1, len(text[1])-5)
			i = 4
			continue
		}
		if i == 4 {
			result.FinalScore = ChangeString(text[1], 1, len(text[1])-5)
			i = 0
		}
		if i == 5 {
			i ++
			continue
		}
		if i == 6 {
			result.TotalScore = ChangeString(text[1], 1, len(text[1])-5)
			i = 0
		}
		if text[1] == "【 平时 】" {
			i = 1
		}
		if text[1] == "【 期末 】" {
			i = 3
		}
		if text[1] == "【 总评 】" {
			i = 5
		}
	}
	fmt.Println(result)
}

func ChangeString(str string, from int, end int) string {
	if len(str) == 0 || end < from {
		return  ""
	}
	tmpStr := []byte(str)
	tmpStr1 := tmpStr[from-1:end-1]
	return string(tmpStr1)
}