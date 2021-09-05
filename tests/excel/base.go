package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

const videoUrl = "https://cn-sxty2-cu-v-02.bilivideo.com/upgcxcode/26/58/370695826/370695826-1-16.mp4?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfq9rVEuxTEnE8L5F6VnEsSTx0vkX8fqJeYTj_lta53NCM=&uipk=5&nbs=1&deadline=1629880700&gen=playurlv2&os=vcache&oi=3733009833&trid=0001aa995399260a4678b8144e7f5105b15eh&platform=html5&upsig=4427073be949b0d72db63f694d795e07&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&cdnid=10924&mid=31701181&bvc=vod&nettype=0&logo=80000000"
const imageUrl = "https://i1.hdslb.com/bfs/archive/449df90131b3a18e2fa2396bddd50e0f6f99a8b6.jpg@640w_400h_1c.webp"

type Root struct {
	ID              int         `json:"id"`
	Root            string         `json:"root"`    // 内容
	Meaning         string         `json:"meaning"` // 意义
	Vocabularies    []Vocabularies `json:"vocabularies"`
	IsLearned       bool           `json:"is_learned"`
	VideoPreviewImg string         `json:"video_preview_img"`
	VideoURL        string         `json:"video_url"`
}

type Vocabularies struct {
	ID      int `json:"id"`
	Word    string `json:"word"`    // 内容
	Meaning string `json:"meaning"` // 词性 + 含义
}

func main() {
	f, err := excelize.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wordCol int
	var rootCol int
	var rootMeaningCol int
	var wordMeaningCol int
	for k, v := range rows[0] {
		switch v {
		case "词汇":
			wordCol = k
		case "核心词根":
			rootCol = k
		case "词根1（词根+意义）":
			rootMeaningCol = k
		case "中文释义":
			wordMeaningCol = k
		}
	}

	fmt.Println(wordCol, rootCol, rootMeaningCol, wordMeaningCol)

	wordCount := 0
	rootCount := 0
	rootMap := make(map[string]*Root)
	for _, row := range rows[1:] {
		wordCount++
		//s := splitWordGender(row[wordMeaningCol])
		word := Vocabularies{
			ID:      wordCount,
			Word:    row[wordCol],
			Meaning: row[wordMeaningCol],
		}
		root, ok := rootMap[row[rootCol]]
		if ok {
			root.Vocabularies = append(root.Vocabularies, word)
		} else {
			rootCount ++
			root = &Root{
				ID:              rootCount,
				Root:            row[rootCol],
				Meaning:         row[rootMeaningCol],
				Vocabularies:    []Vocabularies{word},
				IsLearned:       false,
				VideoPreviewImg: imageUrl,
				VideoURL:        videoUrl,
			}
			rootMap[row[rootCol]] = root
		}
	}

	fmt.Println(rootCount, wordCount)
	roots := make([]*Root, rootCount)
	k := 0
	for _, v := range rootMap {
		roots[k] = v
		k ++
	}

	res, err := marshal(roots)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func marshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func splitWordGender(str string) []string {
	strReplacer := strings.NewReplacer("\t", " ", "\n", " ")
	str = strReplacer.Replace(str)
	genderIndex := strings.Index(str, ".")
	gender := str[:genderIndex+1]
	meaning := str[genderIndex+1:]

	return []string{gender, meaning}
}
