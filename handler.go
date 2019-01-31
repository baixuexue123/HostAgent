package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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
		"%s - %s : %s : %s : %d : %v",
		start, r.Method, r.URL.String(),
		r.RemoteAddr, lrw.statusCode, time.Since(start),
	)
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "HostAgent is serving!\n")
}

func PanicHandler(w http.ResponseWriter, _ *http.Request, rcv interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, http.StatusText(500))
	log.Println("PanicHandler ", rcv)
}

func now(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
}

func info(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	info, _ := host.Info()
	s, _ := json.Marshal(info)
	fmt.Fprint(w, s)
}

func core(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	logical, _ := cpu.Counts(true)
	phys, _ := cpu.Counts(false)
	mapD := map[string]int{"log": logical, "phys": phys}
	s, _ := json.Marshal(mapD)
	fmt.Fprint(w, s)
}

func loadavg(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	loadData, _ := load.Avg()
	s, _ := json.Marshal(loadData)
	fmt.Fprint(w, s)
}

func cpuInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	infoStat, _ := cpu.Info()
	s, _ := json.Marshal(infoStat)
	fmt.Fprint(w, s)
}

func cpuTimes(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	timesStat, _ := cpu.Times(true)
	s, _ := json.Marshal(timesStat)
	fmt.Fprint(w, s)
}

func percpu(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	percent, _ := cpu.Percent(1*time.Second, true)
	s, _ := json.Marshal(percent)
	fmt.Fprint(w, s)
}

func memVir(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	vm, _ := mem.VirtualMemory()
	s, _ := json.Marshal(vm)
	fmt.Fprint(w, s)
}

func memSwap(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	sm, _ := mem.SwapMemory()
	s, _ := json.Marshal(sm)
	fmt.Fprint(w, s)
}

func NewHTTPServer(bind string) *http.Server {
	router := httprouter.New()
	router.PanicHandler = PanicHandler

	router.GET("/", Index)

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
