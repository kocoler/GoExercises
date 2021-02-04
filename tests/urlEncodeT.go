package main

import (
  "fmt"
  "net/url"
)

func main() {
  str := `prefix rdf:<http://www.w3.org/1999/02/22-rdf-syntax-ns#>
prefix rdfs:<http://www.w3.org/2000/01/rdf-schema#>
select distinct ?s ?o1
where {
    ?s rdf:type ?o.
    ?o rdfs:subClassOf ?o1.
}`
  enurl := url.QueryEscape(str)
  fmt.Println(enurl)
}
