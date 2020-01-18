package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	a := "121"
	b := HexToInt([]byte(a))
	fmt.Print("%T",b)
}

func HexToInt(data []byte) int64 {
	//先将字节切片放在Reader中
	//buffer实现了reader，但也直接可以用reader buff := bytes.NewReader(data)
	//bytes.NewReader(data)源码返回了return &Reader{data, 0, -1}
	//而Write(data)源码使用了copy,开辟新空间
	buff := new(bytes.Buffer)
	buff.Write(data)
	var numtemp int64
	//numtemp记得要传指针
	err := binary.Read(buff, binary.BigEndian, &numtemp)

	if err != nil {
		log.Panic(err)
	}
	return  numtemp
}

func toIntSlice(str string) []int {
	var result []int
	tmpStr := []byte(str)
	for i := 0; i <= len(str); i ++ {
		result = append(result,str[1])
	}
}