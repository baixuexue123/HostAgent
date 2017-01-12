package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/shirou/gopsutil/cpu"
	// "github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	// "github.com/shirou/gopsutil/mem"
	// "github.com/shirou/gopsutil/net"
	// "github.com/shirou/gopsutil/process"
)

const __version__ string = "NodeAgent 0.0.1"

func main() {

	router := httprouter.New()

	router.GET("/api/", help)
	router.GET("/api/help/", help)

	router.GET("/api/version/", version)

	router.GET("/api/list/", apiList)

	router.GET("/api/system/", system)
	router.GET("/api/dmidecode/", dmidecode)

	router.GET("/api/now/", now)
	router.GET("/api/uptime/", uptime)

	router.GET("/api/core/", core)
	router.GET("/api/load/", loadavg)
	router.GET("/api/cpu/", version)
	router.GET("/api/percpu/", version)

	router.GET("/api/mem/", version)
	router.GET("/api/mem/used/", version)
	router.GET("/api/memswap/", version)

	router.GET("/api/processlist/", version)
	router.GET("/api/processlist/pid/", version)
	router.GET("/api/processlist/pid/:pid", version)
	router.GET("/api/processcount/", version)

	router.GET("/api/network/", version)
	router.GET("/api/network/interfaces/", version)
	router.GET("/api/network/interface/:iface", version)

	router.GET("/api/hddtemp/", version)

	router.GET("/api/diskio/", version)

	router.GET("/api/fs/", version)

	router.GET("/api/quicklook/", version)

	router.GET("/api/monitors/", version)
	router.GET("/api/monitors/:monitor", version)
	router.PUT("/api/monitors/:monitor", version)

	log.Println("0.0.0.0:9001")
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

func dmidecode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cmd := exec.Command("sudo", "dmidecode")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(w, "The command failed to perform: %s", err)
	}
	fmt.Fprintf(w, "%s", buf)
}

func now(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
}

func uptime(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info, _ := host.Info()
	fmt.Fprintf(w, "%v", info.Uptime)
}

func core(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log, _ := cpu.Counts(true) // logical
	phys, _ := cpu.Counts(false)
	mapD := map[string]int{"log": log, "phys": phys}
	data, _ := json.Marshal(mapD)
	fmt.Fprintf(w, "%s", data)
}

func loadavg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	LoadData, _ := load.Avg()
	fmt.Fprintf(w, "%s", LoadData.String())
}
