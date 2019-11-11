package main
import (

	"io/ioutil"
	"net/http"
)

func main()  {
	url := "https://baike.baidu.com/item/%E5%8D%8E%E4%B8%AD%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6/160405?fr=aladdin"
	reurl, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	re, err := ioutil.ReadAll(reurl.Body)
	if err != nil {
		panic(err)
	}
	defer reurl.Body.Close()
	//fmt.Printf(string(re))
	ioutil.WriteFile("task1.txt",[]byte(re),0644)
}

