package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type APIHandler struct {
	handler http.Handler
}

func NewHandler(router *httprouter.Router) *APIHandler {
	return &APIHandler{router}
}

func (s *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lrw := NewLoggingResponseWriter(w)
	s.handler.ServeHTTP(lrw, r)
	log.Printf(
		"%s  %s  %d  %v %s",
		r.Method, r.URL.String(),
		lrw.statusCode, time.Since(start), r.RemoteAddr,
	)
}

func WriteJSONResponse(w http.ResponseWriter, d interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(d)
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "HostAgent is serving!\n")
}

func PanicHandler(w http.ResponseWriter, _ *http.Request, rcv interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, http.StatusText(500))
	log.Println("PanicHandler ", rcv)
}

func Now(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
}

func HostInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := host.Info()
	WriteJSONResponse(w, data)
}

func HostUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := host.Users()
	WriteJSONResponse(w, data)
}

func HostSensorsTemperatures(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := host.SensorsTemperatures()
	WriteJSONResponse(w, data)
}

func Core(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	logical, _ := cpu.Counts(true)
	phys, _ := cpu.Counts(false)
	data := map[string]int{"log": logical, "phys": phys}
	WriteJSONResponse(w, data)
}

func LoadAvg(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := load.Avg()
	WriteJSONResponse(w, data)
}

func LoadMisc(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := load.Misc()
	WriteJSONResponse(w, data)
}

func CpuInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := cpu.Info()
	WriteJSONResponse(w, data)
}

func CpuTimes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	percpu := qs.Get("percpu") != ""
	data, _ := cpu.Times(percpu)
	WriteJSONResponse(w, data)
}

func CpuPercent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	s := qs.Get("interval")
	if s == "" {
		s = "1"
	}
	interval, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		p, _ := cpu.Percent(time.Duration(interval)*time.Second, true)
		WriteJSONResponse(w, p)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid params interval: %s", s)
	}
}

func MemVm(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := mem.VirtualMemory()
	WriteJSONResponse(w, data)
}

func MemSwap(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := mem.SwapMemory()
	WriteJSONResponse(w, data)
}

func DiskUsageStat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	all := qs.Get("all") != ""
	partitions, _ := disk.Partitions(all)
	data := make(map[string]disk.UsageStat)
	for _, p := range partitions {
		d, _ := disk.Usage(p.Mountpoint)
		data[p.Mountpoint] = *d
	}
	WriteJSONResponse(w, data)
}

func DiskPartitionStat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	all := qs.Get("all") != ""
	partitions, _ := disk.Partitions(all)
	WriteJSONResponse(w, partitions)
}

func DiskIOCountersStat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	all := qs.Get("all") != ""
	partitions, _ := disk.Partitions(all)
	var names []string
	for _, p := range partitions {
		names = append(names, p.Device)
	}
	data, _ := disk.IOCounters(names...)
	WriteJSONResponse(w, data)
}

func NetInterfaces(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	interfaces, _ := net.Interfaces()
	WriteJSONResponse(w, interfaces)
}

func NetIOCounters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	pernic := qs.Get("pernic") != ""
	data, _ := net.IOCounters(pernic)
	WriteJSONResponse(w, data)
}

func NetProtoCounters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	protocols := qs["protocols"]
	data, _ := net.ProtoCounters(protocols)
	WriteJSONResponse(w, data)
}

func NetFilterCounters(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := net.FilterCounters()
	WriteJSONResponse(w, data)
}

func NetConnections(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	kind := qs.Get("kind")
	if kind == "" {
		kind = "all"
	}
	data, _ := net.Connections(kind)
	WriteJSONResponse(w, data)
}

func ProcessPids(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, _ := process.Pids()
	WriteJSONResponse(w, data)
}

type ProcessStat struct {
	Pid           int32   `json:"pid"`
	Name          string  `json:"name"`
	Cwd           string  `json:"cwd"`
	Status        string  `json:"status"`
	Username      string  `json:"username"`
	Cmdline       string  `json:"cmdline"`
	Exe           string  `json:"exe"`
	Terminal      string  `json:"terminal"`
	Uids          []int32 `json:"uids"`
	Gids          []int32 `json:"gids"`
	Background    bool    `json:"background"`
	Foreground    bool    `json:"foreground"`
	Ppid          int32   `json:"ppid"`
	NumThreads    int32   `json:"num_threads"`
	NumFDs        int32   `json:"num_fds"`
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryPercent float32 `json:"mem_percent"`
	CreateTime    int64   `json:"create_time"`
}

func Processes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	processes, _ := process.Processes()
	var data []ProcessStat
	for _, p := range processes {
		proc := ProcessStat{}
		proc.Pid = p.Pid
		proc.Name, _ = p.Name()
		proc.Cwd, _ = p.Cwd()
		proc.Status, _ = p.Status()
		proc.Username, _ = p.Username()
		proc.Uids, _ = p.Uids()
		proc.Gids, _ = p.Gids()
		proc.Cmdline, _ = p.Cmdline()
		proc.Exe, _ = p.Exe()
		proc.Terminal, _ = p.Terminal()
		proc.Ppid, _ = p.Ppid()
		proc.Background, _ = p.Background()
		proc.Foreground, _ = p.Foreground()
		proc.CreateTime, _ = p.CreateTime()
		proc.NumThreads, _ = p.NumThreads()
		proc.NumFDs, _ = p.NumFDs()
		proc.CPUPercent, _ = p.CPUPercent()
		proc.MemoryPercent, _ = p.MemoryPercent()
		data = append(data, proc)
	}
	WriteJSONResponse(w, data)
}

type ProcessStatDetail struct {
	ProcessStat
	MemInfo     *process.MemoryInfoStat  `json:"mem_info"`
	IOCounters  *process.IOCountersStat  `json:"io_counters"`
	Rlimit      []process.RlimitStat     `json:"rlimit"`
	OpenFiles   []process.OpenFilesStat  `json:"open_files"`
	Connections []net.ConnectionStat     `json:"connections"`
	CPUTimes    *cpu.TimesStat           `json:"cpu_times"`
	Threads     map[int32]*cpu.TimesStat `json:"threads"`
}

func ProcessDetail(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	s := params.ByName("pid")
	pid, err := strconv.Atoi(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid params pid: %s", s)
		return
	}
	if ok, _ := process.PidExists(int32(pid)); !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "pid: %s dones not exists", s)
		return
	}
	p, _ := process.NewProcess(int32(pid))
	proc := ProcessStatDetail{}
	proc.Pid = p.Pid
	proc.Name, _ = p.Name()
	proc.Cwd, _ = p.Cwd()
	proc.Status, _ = p.Status()
	proc.Username, _ = p.Username()
	proc.Uids, _ = p.Uids()
	proc.Gids, _ = p.Gids()
	proc.Cmdline, _ = p.Cmdline()
	proc.Exe, _ = p.Exe()
	proc.Terminal, _ = p.Terminal()
	proc.Ppid, _ = p.Ppid()
	proc.Background, _ = p.Background()
	proc.Foreground, _ = p.Foreground()
	proc.CreateTime, _ = p.CreateTime()

	proc.Rlimit, _ = p.Rlimit()
	proc.NumThreads, _ = p.NumThreads()
	proc.NumFDs, _ = p.NumFDs()

	proc.Threads, _ = p.Threads()
	proc.OpenFiles, _ = p.OpenFiles()
	proc.CPUPercent, _ = p.CPUPercent()
	proc.CPUTimes, _ = p.Times()
	proc.MemoryPercent, _ = p.MemoryPercent()
	proc.MemInfo, _ = p.MemoryInfo()
	proc.Connections, _ = p.Connections()
	proc.IOCounters, _ = p.IOCounters()

	WriteJSONResponse(w, proc)
}

func NewHTTPServer(bind string) *http.Server {
	router := httprouter.New()
	router.PanicHandler = PanicHandler

	router.GET("/", Index)
	router.GET("/now/", Now)
	router.GET("/host/info/", HostInfo)
	router.GET("/host/users/", HostUsers)
	router.GET("/host/sensors/", HostSensorsTemperatures)
	router.GET("/core/", Core)
	router.GET("/load/avg/", LoadAvg)
	router.GET("/load/misc/", LoadMisc)
	router.GET("/cpu/info/", CpuInfo)
	router.GET("/cpu/times/", CpuTimes)
	router.GET("/cpu/percent/", CpuPercent)
	router.GET("/mem/vm/", MemVm)
	router.GET("/mem/swap/", MemSwap)
	router.GET("/disk/usage/", DiskUsageStat)
	router.GET("/disk/partitions/", DiskPartitionStat)
	router.GET("/disk/io/", DiskIOCountersStat)
	router.GET("/net/interfaces/", NetInterfaces)
	router.GET("/net/io/counters/", NetIOCounters)
	router.GET("/net/proto/counters/", NetProtoCounters)
	router.GET("/net/filter/counters/", NetFilterCounters)
	router.GET("/net/connections/", NetConnections)
	router.GET("/processes/", Processes)
	router.GET("/process/:pid/", ProcessDetail)
	router.GET("/pids/", ProcessPids)

	server := &http.Server{
		Addr:    bind,
		Handler: NewHandler(router),
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		server.Shutdown(ctx)
	}()

	return server
}
