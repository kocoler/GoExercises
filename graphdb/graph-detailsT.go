package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// http://localhost:7200/rest/explore-graph/node
var NodeDetailsURL = "http://localhost:7200/rest/explore-graph/properties"

// GetNodeDetails get node details.
func GetNodeDetails(node string) ([]byte, error) {
	method := "GET"
	data := url.Values{}
	data.Set("config", "default")
	data.Set("includeInferred", "false")
	data.Set("iri", node)
	data.Set("sameAsState", "false")
	client := &http.Client{}

	// fmt.Println(strings.NewReader(data.Encode()))
	req, err := http.NewRequest(method, "http://localhost:7200/rest/explore-graph/properties?config=default&includeInferred=false&sameAsState=false&iri=http:%2F%2Fid.nlm.nih.gov%2Fmesh%2FD001172", strings.NewReader(data.Encode()))
	if err != nil {
		return []byte{}, err
	}
	fmt.Println(req.RequestURI, req.Form)

	req.Header.Add("X-GraphDB-Repository", "NLPMicrobeKG")

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	// fmt.Println(string(body))

	return body, nil
}

func main() {
	body, _ := GetNodeDetails("http://id.nlm.nih.gov/mesh/D001172")
	var res map[string]interface{}
	err := json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}
