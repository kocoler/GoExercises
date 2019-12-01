package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Product _
type Product struct {
	Name      []string  `json:"name"`
	ProductID int64   `json:"product_id,omitempty"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,omitempty"`
}

func main() {
	p := &Product{}
	testStrings := []string{
		"I am string1",
		"I am string2",
		"I am Art3mis.",
		"I am Parzival.",
	}
	p.Name = testStrings
	p.IsOnSale = false
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0
	a,_ := json.Marshal(p)
	fmt.Println(string(a))
	d := `{"name":["I am string1","I am string2","I am Art3mis.","I am Parzival."],"number":10000,"price":2499}`
	e := &Product{}
	err := json.Unmarshal([]byte(d),e)
	log.Println(err)
	for _,v := range p.Name {
		fmt.Println(v)
	}
	//fmt.Println(p)
}