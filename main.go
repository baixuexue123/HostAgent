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
	"github.com/shirou/gopsutil/mem"
	// "github.com/shirou/gopsutil/net"
	// "github.com/shirou/gopsutil/process"
)

const Version = "NodeAgent 0.0.1"

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
	router.GET("/api/cpu/info/", cpuInfo)
	router.GET("/api/cpu/times/", cpuTimes)
	router.GET("/api/percpu/", percpu)

	router.GET("/api/mem/", memVir)
	router.GET("/api/memswap/", memSwap)

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
	log.Fatal(http.ListenAndServe(":9001", router))
}

func version(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, Version)
}

func help(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "NodeAgent Help")
}

func apiList(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "NodeAgent ApiList")
}

func system(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "system")
}

func dmidecode(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	cmd := exec.Command("sudo", "dmidecode")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(w, "The command failed to perform: %s", err)
	}
	fmt.Fprintf(w, "%s", buf)
}

func now(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
}

func uptime(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	info, _ := host.Info()
	fmt.Fprintf(w, "%v", info.Uptime)
}

func core(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	logical, _ := cpu.Counts(true)
	phys, _ := cpu.Counts(false)
	mapD := map[string]int{"log": logical, "phys": phys}
	data, _ := json.Marshal(mapD)
	fmt.Fprintf(w, "%s", data)
}

func loadavg(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	loadData, _ := load.Avg()
	fmt.Fprintf(w, "%s", loadData.String())
}

func cpuInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	infoStat, _ := cpu.Info()
	data, _ := json.Marshal(infoStat)
	fmt.Fprintf(w, "%s", data)
}

func cpuTimes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	timesStat, _ := cpu.Times(true)
	data, _ := json.Marshal(timesStat)
	fmt.Fprintf(w, "%s", data)
}

func percpu(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	percent, _ := cpu.Percent(1*time.Second, true)
	data, _ := json.Marshal(percent)
	fmt.Fprintf(w, "%s", data)
}

func memVir(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	vm, _ := mem.VirtualMemory()
	fmt.Fprintf(w, "%s", vm.String())
}

func memSwap(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	sm, _ := mem.SwapMemory()
	fmt.Fprintf(w, "%s", sm.String())
}
