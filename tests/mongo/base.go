package main

import (
	"context"
	"fmt"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mongoClient *mongo.Client

func init() {
	getClient()
}

func getClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	mongoClient = client
}

type Root struct {
	//ID string `bson:"_id" json:"_id"`
	Content   string `bson:"content" json:"content"`     // 内容
	Sense     string `bson:"sense" json:"sense"`         // 词根+意义
	Additions string `bson:"additions" json:"additions"` // 额外字段
}

type Word struct {
	//ID string `bson:"_id" json:"_id"`
	RootID1      string `bson:"rootid1" json:"rootid1"`             // 词根1
	RootID2      string `bson:"rootid2" json:"rootid2"`             // 词根2
	RootID3      string `bson:"rootid3" json:"rootid3"`             // 词根3
	Content      string `bson:"content" json:"content"`             // 单词内容
	Sense        string `bson:"sense" json:"sense"`                 // 单词解释 + 词性
	Soundmark    string `bson:"soundmark" json:"soundmark"`         // 音标
	Paraphrase   string `bson:"paraphrase" json:"paraphrase"`       // 词根释义
	ParaphraseEn string `bson:"paraphrase_en" json:"paraphrase_en"` // 英文释义
	ParaphraseZh string `bson:"paraphrase_zh" json:"paraphrase_zh"` // 中文释义
	Derivative   string `bson:"derivative" json:"derivative"`       // 派生词
	Additions    string `bson:"additions" json:"additions"`
}

func insert(ctx context.Context, rootsCollection mongo.Collection) {
	root1 := Root{
		Content:      "aud(it)",
		Sense:        "ag＝to do 做",
		Additions:    "",
	}

	res, err := rootsCollection.InsertOne(ctx, root1)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.InsertedID, err)
}

func main() {
	ctx := context.Background()
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	rootsCollection := mongoClient.Database("sourceword").Collection("roots")


}
