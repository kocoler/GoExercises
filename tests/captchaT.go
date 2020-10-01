package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	id := captcha.NewLen(5)
	//fmt.Print(id)
	//captcha.
	//captcha.
	f, err := os.Create("w.png")
	if err != nil {
		fmt.Println(err)
	}
	//digits := captcha.RandomDigits(5)
	err = captcha.WriteImage(f, id, 320, 80)
	if err != nil {
		fmt.Println(err)
	}
	//SaveImage("ww.png", newImage)
	var digits string
	fmt.Scanf("%s", &digits)
	res := captcha.VerifyString(id, digits)
	//res2 := captcha.Verify(id, []byte{'2', 's', 's', 's', 'w'})
	//fmt.Println()
	fmt.Println(res, digits)
	//fmt.Println(res, res2)
}

func SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(p)
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}
