package main

import (
	"context"
	"fmt"
	elastic "github.com/olivere/elastic"
)

func main() {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	

}