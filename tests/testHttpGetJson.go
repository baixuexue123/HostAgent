package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("search query failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		resp.Body.Close()
		return err
	}

	resp.Body.Close()

	return nil
}

type JsonData struct {
	code int
	msg  string
}

func main() {
	foo1 := new(JsonData)
	getJson("http://localhost:8888/string", foo1)
	fmt.Println(foo1)

	foo2 := JsonData{}
	getJson("http://localhost:8888/json", &foo2)
	fmt.Println(foo2.msg, foo2.msg)
}
