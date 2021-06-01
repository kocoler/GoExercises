package main

import (
	"bytes"
	"log"

	"github.com/beyondstorage/go-service-fs/v3"
	"github.com/beyondstorage/go-storage/v4/pairs"
)

func main() {
	// Init a service.
	store, err := fs.NewStorager(pairs.WithWorkDir("tmp"))
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

	content := []byte("Hello, world!")
	length := int64(len(content))
	r := bytes.NewReader(content)

	_, err = store.Write("hello", r, length)
	if err != nil {
		log.Fatalf("write failed: %v", err)
	}

	var buf bytes.Buffer

	_, err = store.Read("hello", &buf)
	if err != nil {
		log.Fatalf("storager read: %v", err)
	}

	log.Printf("%s", buf.String())
}
