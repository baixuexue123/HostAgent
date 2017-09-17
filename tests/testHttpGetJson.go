package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JsonData struct {
	Code int    `json: "code"`
	Msg  string `json: "msg"`
}

func getJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	data1 := new(JsonData)
	if err := getJson("http://localhost:8888/json", data1); err != nil {
		log.Fatal(err)
	}
	fmt.Println(data1)
	fmt.Println(data1.Code)
	fmt.Printf("msg: %s\n", data1.Msg)

	data2 := JsonData{}
	if err := getJson("http://localhost:8888/json", &data2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data2)
	fmt.Println(data2.Code)
	fmt.Printf("msg: %s\n", data2.Msg)
}
