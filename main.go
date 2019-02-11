package main

import (
	"flag"
	"log"
	"os"
)

var Args struct {
	Bind  string
	Debug bool
}

func ParseCli() {
	flag.StringVar(&Args.Bind, "bind", "127.0.0.1:9090", "addr:port")
	flag.BoolVar(&Args.Debug, "debug", true, "debug")
	flag.Parse()
}

func main() {
	ParseCli()
	s := NewHTTPServer(Args.Bind)
	log.SetOutput(os.Stdout)
	log.Println("HostAgent is started ", Args.Bind)
	log.Fatal(s.ListenAndServe())
}
