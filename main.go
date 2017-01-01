package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var __version__ = "0.0.1"

func main() {

	router := httprouter.New()

	router.GET("/api/version/", version)

	router.GET("/api/help/", help)

	router.GET("/api/list/", apiList)

	router.GET("/api/system/", counter)

	router.GET("/api/now/", counter)
	router.GET("/api/uptime/", counter)

	router.GET("/api/core/", counter)
	router.GET("/api/load/", counter)
	router.GET("/api/cpu/", counter)
	router.GET("/api/percpu/", counter)

	router.GET("/api/mem/", counter)
	router.GET("/api/mem/used/", counter)
	router.GET("/api/memswap/", counter)

	router.GET("/api/processlist/", counter)
	router.GET("/api/processlist/pid/", counter)
	router.GET("/api/processlist/pid/:pid", counter)
	router.GET("/api/processcount/", counter)

	router.GET("/api/network/", counter)
	router.GET("/api/network/interfaces/", counter)
	router.GET("/api/network/interface/:iface", counter)

	router.GET("/api/hddtemp/", counter)

	router.GET("/api/diskio/", counter)

	router.GET("/api/fs/", counter)

	router.GET("/api/quicklook/", counter)

	router.GET("/api/monitors/", counter)
	router.GET("/api/monitors/:monitor", counter)
	router.PUT("/api/monitors/:monitor", counter)

	log.Fatal(http.ListenAndServe("0.0.0.0:9001", router))
}

func version(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s", __version__)
}

func help(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "NodeAgent Help")
}

func apiList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "NodeAgent ApiList")
}

func system(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "system")
}
