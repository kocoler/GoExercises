package main

import (
	"flag"
	"fmt"
	"os"
)

var s = flag.String("s", "", "Your sid.")
var p = flag.String("p", "", "Your password.")
var d = flag.Bool("d", false, "Whether output information.")

func main() {
	for i, args := range os.Args {
		fmt.Printf("args[%d]=%s\n",i,args)
	}

	fmt.Println(len(os.Args))

	flag.Parse()

	fmt.Println(*p, *d, *s)
}
