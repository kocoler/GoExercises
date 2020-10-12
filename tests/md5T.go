package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

func main() {
	data := []byte("alice")
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	log.Println(md5str1)
}
