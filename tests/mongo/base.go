package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//type Root struct {
//	//ID string `bson:"_id" json:"_id"`
//	Content   string `bson:"content" json:"content"`     // 内容
//	Sense     string `bson:"sense" json:"sense"`         // 词根+意义
//	Additions string `bson:"additions" json:"additions"` // 额外字段
//}

type Root struct {
	ID              int    `json:"id"`
	Root            string `json:"root"`    // 内容
	Meaning         string `json:"meaning"` // 意义
	IsLearned       bool   `json:"is_learned"`
	VideoPreviewImg string `json:"video_preview_img"`
	VideoURL        string `json:"video_url"`
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

func insert(ctx context.Context, rootsCollection mongo.Collection, value interface{}) {
	//root1 := Root{
	//	ID:              0,
	//	Root:            "",
	//	Meaning:         "",
	//	IsLearned:       false,
	//	VideoPreviewImg: "",
	//	VideoURL:        "",
	//}

	res, err := rootsCollection.InsertOne(ctx, value)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.InsertedID, err)
}

func queryOne(ctx context.Context, rootsCollection mongo.Collection) {
	queryR := rootsCollection.FindOne(context.TODO(), bson.D{bson.E{Key: "user_id", Value: 1}})

	bytesData, err := queryR.DecodeBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println(bytesData.Lookup("roots_status").StringValue())
}

func query(ctx context.Context, rootsCollection mongo.Collection) {
	opts := options.Find()//.SetSort(bson.D{{"id", 1}}).SetSkip(0).SetLimit(10)
	// cursor, err := rootsCollection.Find(context.TODO(), bson.D{bson.E{Key: "root", Value: bson.D{{"$regex", "a"}}}}, opts)
	cursor, err := rootsCollection.Find(context.TODO(), bson.D{bson.E{Key: "user_id", Value: 1}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, v := range results {
		fmt.Println(v)
	}
}

func main() {
	ctx := context.Background()
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	rootsCollection := mongoClient.Database("sourceword").Collection("user_roots")

	queryOne(ctx, *rootsCollection)
}
