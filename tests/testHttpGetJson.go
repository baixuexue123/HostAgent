package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJson(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// return json.NewDecoder(res.Body).Decode(target)
	return json.Unmarshal(res.Body, target)
}

type JsonData struct {
	code int
	msg  string
}

func main() {
	foo1 := new(JsonData)
	getJson("http://localhost:8888/json", foo1)
	fmt.Println(foo1)

	foo2 := JsonData{}
	getJson("http://localhost:8888/json", &foo2)
	fmt.Println(foo2.msg, foo2.msg)
}
