package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pwd,_ := os.Getwd()

	fileInfo, err := ioutil.ReadDir(pwd+"/github.com/casbin/casbin-forum/.git/refs/heads")
	for a, v := range fileInfo {
		if v.Name() == "master" {
			fmt.Println(a, v.ModTime())
		}
	}
	fmt.Println(fileInfo[1].ModTime())

	return
	//获取当前目录下的文件或目录名(包含路径)
	file, err := os.Open(pwd+"/github.com/casbin/casbin-forum/.git/refs/heads/master")
	defer file.Close()
	//reader := bufio.NewReader(file)
	//fmt.Println(reader)
	content, err := ioutil.ReadFile(pwd+"/github.com/casbin/casbin-forum/.git/refs/heads/master")
	fmt.Println(string(content[:7]), err)
	//log.Println(file., err)
	return
	filepathNames,err := filepath.Glob(filepath.Join(pwd,"*"))
	if err != nil {
		log.Fatal(err)
	}

	for i := range filepathNames {
		fmt.Println(filepathNames[i]) //打印path
	}
}
