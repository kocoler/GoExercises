package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const videoUrl = "https://cn-sxty2-cu-v-02.bilivideo.com/upgcxcode/26/58/370695826/370695826-1-16.mp4?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfq9rVEuxTEnE8L5F6VnEsSTx0vkX8fqJeYTj_lta53NCM=&uipk=5&nbs=1&deadline=1629880700&gen=playurlv2&os=vcache&oi=3733009833&trid=0001aa995399260a4678b8144e7f5105b15eh&platform=html5&upsig=4427073be949b0d72db63f694d795e07&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&cdnid=10924&mid=31701181&bvc=vod&nettype=0&logo=80000000"
const imageUrl = "https://i1.hdslb.com/bfs/archive/449df90131b3a18e2fa2396bddd50e0f6f99a8b6.jpg@640w_400h_1c.webp"

type videoInfo struct {
	videoUrl string
	imageUrl string
}

var m = map[string]videoInfo {"": {videoUrl: "", imageUrl: ""}}

var mongoClient *mongo.Client

func init() {
	getClient()
}

func getClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: "",
		Password: "",
	}

	// mongodb://112.126.76.187:27017
	client, err := mongo.Connect(ctx, options.Client().SetAuth(credential).ApplyURI("mongodb://admin:@112.126.76.187:27017/?ssl=false&authMechanism=SCRAM-SHA-1"))
	if err != nil {
		panic(err)
	}

	mongoClient = client
}

type Root struct {
	ID              int    `bson:"id" json:"id"`
	Root            string `bson:"root" json:"root"`                 // 内容
	Meaning         string `bson:"meaning" json:"meaning"`           // 意义
	Vocabularies    []int  `bson:"vocabularies" json:"vocabularies"` // 所有相关单词
	VideoPreviewImg string `bson:"videopreviewimg" json:"video_preview_img"`
	VideoURL        string `bson:"videourl" json:"video_url"`
}

type Derivative struct {
	Word      string `bson:"word"`       // 内容
	MeaningZh string `bson:"meaning_zh"` // 中文释义
}

type Word struct {
	ID                 int          `bson:"id"`
	StructuralAnalysis string       `bson:"structural_analysis"` // 结构分析
	Word               string       `bson:"word"`                // 内容
	MeaningEn          string       `bson:"meaning_en"`          // 英文释义
	MeaningZh          string       `bson:"meaning_zh"`          // 中文释义
	ExampleSentences   string       `bson:"example_sentences"`   // 例句
	ExamGrading        []int        `bson:"exam_grading"`        // 考试分级
	Derivatives        []Derivative `bson:"derivative" `         // 派生词
}

var words []*Word
var roots []*Root

func main() {
	importRootsAndWords()
}

func importRootsAndWords() {
	f, err := excelize.OpenFile("词库.xlsx")
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
	var wordMeaningZhCol int
	var wordMeaningEnCol int
	var StructuralAnalysisCol int
	var ExampleSentencesCol int
	var ExampleGradingFromCol int
	var ExampleGradingEndCol int
	var Derivatives1Col int
	var Derivatives2Col int
	for k, v := range rows[0] {
		switch v {
		case "词汇":
			wordCol = k
		case "结构分析":
			StructuralAnalysisCol = k
		case "例句":
			ExampleSentencesCol = k
		case "CET4":
			ExampleGradingFromCol = k
		case "GRE":
			ExampleGradingEndCol = k
		case "派生词1":
			Derivatives1Col = k
		case "派生词2":
			Derivatives2Col = k
		case "核心词根":
			rootCol = k
		case "词素1（词素+意义）":
			rootMeaningCol = k
		case "中文释义":
			wordMeaningZhCol = k
		case "英文释义":
			wordMeaningEnCol = k
		}
	}

	fmt.Println(wordCol, rootCol, rootMeaningCol, wordMeaningZhCol, ExampleGradingFromCol, Derivatives2Col)
	//return
	wordCount := 0
	rootCount := 0
	rootMap := make(map[string]*Root)
	//fmt.Println(len(rows))
	//return
	for k, row := range rows[1:] {
		if len(row) > 1 {
			logs := fmt.Sprintf("importing ... col: %d, word: %s", k, row[wordCol])
			fmt.Println(logs)
		}

		//fmt.Println(len(row))
		//continue
		if len(row) < wordMeaningZhCol {
			continue
		}
		wordCount++
		//s := splitWordGender(row[wordMeaningCol])
		exampleGrading := make([]int, ExampleGradingEndCol-ExampleGradingFromCol+1)
		for i := ExampleGradingFromCol; i <= ExampleGradingEndCol && i < len(row); i++ {
			if row[i] == "1.0" {
				exampleGrading[i-ExampleGradingFromCol] = 1
			}
		}
		derivative := make([]Derivative, 2)

		if len(row) > Derivatives1Col {
			derivative[0] = Derivative{Word: row[Derivatives1Col], MeaningZh: row[Derivatives1Col+1]}
		}
		if len(row) > Derivatives2Col {
			derivative[1] = Derivative{Word: row[Derivatives2Col], MeaningZh: row[Derivatives2Col+1]}
		}

		word := Word{
			ID:                 wordCount,
			StructuralAnalysis: row[StructuralAnalysisCol],
			Word:               row[wordCol],
			MeaningEn:          row[wordMeaningEnCol],
			MeaningZh:          row[wordMeaningZhCol],
			ExamGrading:        exampleGrading,
			Derivatives:        derivative,
		}
		if ExampleSentencesCol < len(row) {
			word.ExampleSentences = row[ExampleSentencesCol]
		}
		words = append(words, &word)
		root, ok := rootMap[row[rootCol]]
		if ok {
			root.Vocabularies = append(root.Vocabularies, wordCount)
		} else {
			rootCount++
			root = &Root{
				ID:              rootCount,
				Root:            row[rootCol],
				Meaning:         row[rootMeaningCol],
				Vocabularies:    []int{wordCount},
				VideoPreviewImg: imageUrl,
				VideoURL:        videoUrl,
			}
			rootMap[row[rootCol]] = root
		}
	}

	fmt.Println(rootCount, wordCount)

	roots = make([]*Root, rootCount)
	k := 0
	for _, v := range rootMap {
		roots[k] = v
		k++
	}

	fmt.Println(k)

	for _, v := range roots {
		_, err := mongoClient.Database("sourceword").Collection("roots").InsertOne(context.TODO(), v)
		if err != nil {
			panic(err)
		}
		//fmt.Println(res.InsertedID, err)
	}
	for _, v := range words {
		_, err := mongoClient.Database("sourceword").Collection("words").InsertOne(context.TODO(), v)
		if err != nil {
			panic(err)
		}
		//fmt.Println(res.InsertedID, err)
	}
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
