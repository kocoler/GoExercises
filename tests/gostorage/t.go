package main

import (
	"fmt"
	"log"

	"github.com/beyondstorage/go-service-fs/v3"
	"github.com/beyondstorage/go-storage/v4/pairs"
)

func main() {
	// Init a service.
	store, err := fs.NewStorager(pairs.WithWorkDir("./"))
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

	//content := []byte("Hello, world!")
	//length := int64(len(content))
	//r := bytes.NewReader(content)
	//
	//_, err = store.Write("hello2", r, length)
	//if err != nil {
	//	log.Fatalf("write failed: %v", err)
	//}
	//
	//var buf bytes.Buffer
	//
	//_, err = store.Read("hello2", &buf)
	//if err != nil {
	//	log.Fatalf("storager read: %v", err)
	//}
	fmt.Println()
	it, err := store.List("/Users/mac/go/src/github.com/kocoler/GoExercises/tests/gostorage/tmp", pairs.WithContinuationToken("hello3"))
	if err != nil {
		fmt.Println(err)
	}
	//log.Printf("%s", buf.String())
	obj, err := it.Next()
	for obj != nil {
		fmt.Println("obj.Path:", obj.Path)
		fmt.Println("obj.Mode:", obj.Mode)
		fmt.Println("obj.Id", obj.ID)
		fmt.Println("continue token", it.ContinuationToken())
		obj, err = it.Next()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(obj, it)
}
