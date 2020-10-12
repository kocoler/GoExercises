package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//log.Fatal()
	db, err := sql.Open("mysql","root:CCNU200181_dx@tcp(localhost:3306)/ginUsers")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(1)
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id=129").Scan(&name)
	//log.Println(err)
	//fmt.Println(name)
	db.Exec("alter table users add column aa varchar(128)")
}

