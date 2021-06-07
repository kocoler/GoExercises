package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Clean("/wew/"))
	fmt.Println(filepath.ToSlash("/www/www"))
	fmt.Println("base:", path.Base("/www/www/"))
	s := "/www/ww/"
	fmt.Println(s[1:], s[:len(s) - 1])
}
