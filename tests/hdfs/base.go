package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"io/ioutil"
)

func main() {
	client, err := hdfs.New("localhost:9000")
	if err != nil {
		fmt.Println("open error: ", err)
	}

	//file, err := client.Open("/testfile/wwww.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	writer, err := client.Create("/testfile/ww.txt")
	if err != nil {
		fmt.Println(err)
	}
	n, err := writer.Write([]byte("www"))
	fmt.Println("write!", n, err)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader, err := client.Open("/testfile/ww.txt")
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("readall error:", err)
	}
	fmt.Println(string(bytes))
	// => Abominable are the tumblers into which he pours his poison.
}

