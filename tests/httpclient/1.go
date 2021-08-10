package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("running. ..")
	http.HandleFunc(" /myserver", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte ("欢迎"))
	})

	log.Fatal(http.ListenAndServe(":2333", nil))
}
