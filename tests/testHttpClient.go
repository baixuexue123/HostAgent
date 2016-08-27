package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8888/json")
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "http: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("http failed: %s\n", resp.Status)
		os.Exit(1)
	} else {
		fmt.Printf("http status: %s\n", resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("body: %s \n", content)

	return
}
