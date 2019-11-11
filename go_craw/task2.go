package main

import (
	"github.com/goquery"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	reurl := "https://studygolang.com/topics"
	re, err := http.Get(reurl)
	if err != nil {
		panic(err)
	}
	dom, err := goquery.NewDocumentFromReader(re.Body)
	if err != nil {
		panic(err)
	}
	defer re.Body.Close()
	var x int
	fmt.Printf("输入x:")
	fmt.Scanf("%d",&x)
	x --
	dom.Find(".title").Each(func(i int, selection *goquery.Selection) {
		/*selection.Find(".title").Each(func(i int, selection1 *goquery.Selection) {
		//	fmt.Println(selection1.Text())
		})*/
		if i == x {
			str := strings.Replace(selection.Text()," ","",-1)
			fmt.Println(str)
		}

	})


}

