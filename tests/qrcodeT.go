package main

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	suffix := "act"
	name := `2019级寝室卫生清理及安全检查劳育活动`
	id := `335`
	from := `2020-12-20 19:00:00`
	to := `2020-12-20 20:00:00`
	place := `学生寝室`
	contentSlice := []string{suffix, name, id, from, to, place}

	content := strings.Join(contentSlice, "%")

	err := qrcode.WriteColorFile(content, qrcode.Medium, 256, color.White, color.Black, "qr.png")
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Done!")
}
