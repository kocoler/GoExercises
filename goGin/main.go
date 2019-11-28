package main

import (
	"goGin/model"
	"goGin/router"
	"log"
)
func main()  {
	model.Db.Init()
	defer model.Db.Close()
	router.Init()
	if err := router.Router.Run(":9090"); err!= nil {
		log.Fatal(err)
	}
	log.Println("SUCCEED IN SILENCE")
}