package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Where struct {
	Field []string `json:"field"`
	Item  []string `json:"item"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")

	fmt.Println(r.RequestURI)

	queryValues := r.URL.Query()

	s := queryValues.Get("gran")
	fmt.Printf("gran: %s", s)
	s = queryValues.Get("select")
	fmt.Printf("select: %s", s)

	Select := make(map[string]Where)
	if err := json.Unmarshal([]byte(s), &Select); err != nil {
		fmt.Fprintf(w, "select: %v is not json", s)
		return
	}
	for Type, where := range Select {
		fmt.Println(Type)
		fmt.Println(where)
	}

	_, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(w, "gran: %v is not int", s)
		return
	}

}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	fmt.Fprintf(w, "hello, %v\n", ps)
	fmt.Fprintf(w, "%v\n", r)
}

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s; %s; %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	for k, v := range r.PostForm {
		fmt.Fprintf(w, "PostForm[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Params: %v\n", ps)

	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "body: %s\n", content)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/handler", handler)
	router.POST("/handler", handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
