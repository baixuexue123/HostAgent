package main

import (
    "encoding/json"
    "fmt"
    "net/url"
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
)


// Client-Get
func main() {
    u, _ := url.Parse("http://localhost:9001/xiaoyue")
    q := u.Query()
    q.Set("username", "user")
    q.Set("password", "passwd")
    u.RawQuery = q.Encode()
    res, err := http.Get(u.String());
    if err != nil {
        log.Fatal(err)
        return
    }
    result, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err) return
    }
    fmt.Printf("%s", result)
}

// Client-Post
type Server struct {
    ServerName string
    ServerIP   string
}

type ServerSlice struct {
    Servers []Server
    ServersID  string
}


func main() {
    var s ServerSlice
    var newServer Server
    newServer.ServerName = "Guangzhou_VPN"
    newServer.ServerIP = "127.0.0.1"
    s.Servers = append(s.Servers, newServer)
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.2"})
    s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.3"})

    s.ServersID = "team1"

    b, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }

    body := bytes.NewBuffer([]byte(b))
    res,err := http.Post("http://localhost:9001/xiaoyue", "application/json;charset=utf-8", body)
    if err != nil {
        log.Fatal(err)
        return
    }
    result, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
        return
    }
    fmt.Printf("%s", result)
}


// *******************************************************

func main() {
    url := "http://restapi3.apiary.io/notes"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

// *******************************************************

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
    Bar string
}

func main() {
    foo1 := new(Foo) // or &Foo{}
    getJson("http://example.com", foo1)
    println(foo1.Bar)

    // alternately:

    foo2 := Foo{}
    getJson("http://example.com", &foo2)
    println(foo2.Bar)
}

// *******************************************************

func main() {

    data := map[string]interface{}{}

    r, _ := http.Get("http://api.stackoverflow.com/1.1/tags?pagesize=100&page=1")
    defer r.Body.Close()

    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &data)

    fmt.Println("Total:", data["total"], "page:", data["page"], "pagesize:", data["pagesize"])
    // Total: 34055 page: 1 pagesize: 100
}

// *******************************************************

func test(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t test_struct   
    err := decoder.Decode(&t)
    if err != nil {
        panic()
    }
    log.Println(t.Test)
}
