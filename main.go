package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	http.HandleFunc("/api/list/", handler)

	http.HandleFunc("/api/version/", counter)

	http.HandleFunc("/api/help/", counter)

	http.HandleFunc("/api/system/", counter)

	http.HandleFunc("/api/now/", counter)
	http.HandleFunc("/api/uptime/", counter)

	http.HandleFunc("/api/core/", counter)
	http.HandleFunc("/api/load/", counter)
	http.HandleFunc("/api/cpu/", counter)
	http.HandleFunc("/api/percpu/", counter)

	http.HandleFunc("/api/mem/", counter)
	http.HandleFunc("/api/mem/used/", counter)
	http.HandleFunc("/api/memswap/", counter)

	http.HandleFunc("/api/processlist/", counter)
	http.HandleFunc("/api/processlist/pid/", counter)
	http.HandleFunc("/api/processlist/pid/:pid", counter)
	http.HandleFunc("/api/processcount/", counter)

	http.HandleFunc("/api/network/", counter)
	http.HandleFunc("/api/network/interfaces/", counter)
	http.HandleFunc("/api/network/interface/:iface", counter)

	http.HandleFunc("/api/hddtemp/", counter)

	http.HandleFunc("/api/diskio/", counter)

	http.HandleFunc("/api/fs/", counter)

	http.HandleFunc("/api/monitor/", counter)

	http.HandleFunc("/api/quicklook/", counter)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count %d\n", 1000)
}
