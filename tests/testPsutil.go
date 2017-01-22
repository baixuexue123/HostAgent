package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"

	"../humanize"
)

func testHostInfo() {
	n, _ := host.Info()
	fmt.Printf("Hostname : %v\n", n.Hostname)
	fmt.Printf("Uptime : %v\n", n.Uptime)
	fmt.Printf("BootTime : %v\n", n.BootTime)
	fmt.Printf("Procs : %v\n", n.Procs)
	fmt.Printf("OS : %s\n", n.OS)
	fmt.Printf("Platform : %v\n", n.Platform)
	fmt.Printf("PlatformFamily : %v\n", n.PlatformFamily)
	fmt.Printf("PlatformVersion : %v\n", n.PlatformVersion)
	fmt.Printf("KernelVersion : %v\n", n.KernelVersion)
	fmt.Printf("VirtualizationSystem : %v\n", n.VirtualizationSystem)
	fmt.Printf("VirtualizationRole : %v\n", n.VirtualizationRole)
	fmt.Printf("HostID : %v\n", n.HostID)
}

func testUserStat() {
	users, _ := host.Users()
	fmt.Println(len(users))
	for _, u := range users {
		fmt.Println()
		fmt.Printf("User : %s\n", u.User)
		fmt.Printf("Terminal : %s\n", u.Terminal)
		fmt.Printf("Host : %s\n", u.Host)
		fmt.Printf("Started : %v\n", u.Started)
	}
}

func testCpuTimesStat() {
	timesStat, _ := cpu.Times(true)

	// fmt.Printf("TimesStat: %v", timesStat)

	for _, stat := range timesStat {
		fmt.Println()
		fmt.Printf("Total : %v\n", stat.Total())
		fmt.Printf("CPU : %s\n", stat.CPU)
		fmt.Printf("User : %v\n", stat.User)
		fmt.Printf("System : %v\n", stat.System)
		fmt.Printf("Idle : %v\n", stat.Idle)
		fmt.Printf("Nice : %v\n", stat.Nice)
		fmt.Printf("Iowait : %v\n", stat.Iowait)
		fmt.Printf("Irq : %v\n", stat.Irq)
		fmt.Printf("Softirq : %v\n", stat.Softirq)
		fmt.Printf("Steal : %v\n", stat.Steal)
		fmt.Printf("Guest : %v\n", stat.Guest)
		fmt.Printf("GuestNice : %v\n", stat.GuestNice)
		fmt.Printf("Stolen : %v\n", stat.Stolen)
	}
}

func testCpuInfo() {
	count, _ := cpu.Counts(true) // logical
	fmt.Printf("Cpu Counts: %v\n", count)

	percent, _ := cpu.Percent(1*time.Second, false)
	fmt.Printf("Cpu Percent: %v\n", percent)

	infoStat, _ := cpu.Info()
	for _, stat := range infoStat {
		fmt.Println()
		fmt.Printf("CPU : %v\n", stat.CPU)
		fmt.Printf("VendorID : %s\n", stat.VendorID)
		fmt.Printf("Family : %s\n", stat.Family)
		fmt.Printf("Model : %s\n", stat.Model)
		fmt.Printf("Stepping : %v\n", stat.Stepping)
		fmt.Printf("PhysicalID : %s\n", stat.PhysicalID)
		fmt.Printf("CoreID : %s\n", stat.CoreID)
		fmt.Printf("Cores : %v\n", stat.Cores)
		fmt.Printf("ModelName : %s\n", stat.ModelName)
		fmt.Printf("Mhz : %v\n", stat.Mhz)
		fmt.Printf("CacheSize : %v\n", stat.CacheSize)
		fmt.Printf("Flags : %v\n", stat.Flags)
	}
}

func testLoadavg() {
	loadavg, _ := load.Avg()

	fmt.Printf("Load1 : %v\n", loadavg.Load1)
	fmt.Printf("Load5 : %v\n", loadavg.Load5)
	fmt.Printf("Load15 : %v\n", loadavg.Load15)
}

func testMem() {
	// func VirtualMemory() (*VirtualMemoryStat, error)
	vm, _ := mem.VirtualMemory()

	fmt.Println("------ VirtualMemory ------")
	fmt.Printf("Total: %v\n", vm.Total)
	fmt.Printf("Available: %v\n", vm.Available)
	fmt.Printf("Used: %v\n", vm.Used)
	fmt.Printf("UsedPercent: %f%%\n", vm.UsedPercent)
	fmt.Printf("Free: %v\n", vm.Free)

	fmt.Printf("Active: %v\n", vm.Active)
	fmt.Printf("Inactive: %v\n", vm.Inactive)
	fmt.Printf("Wired: %v\n", vm.Wired)

	fmt.Printf("Buffers: %v\n", vm.Buffers)
	fmt.Printf("Cached: %v\n", vm.Cached)
	fmt.Printf("Writeback: %v\n", vm.Writeback)
	fmt.Printf("Dirty: %v\n", vm.Dirty)
	fmt.Printf("WritebackTmp: %v\n", vm.WritebackTmp)

	// func SwapMemory() (*SwapMemoryStat, error)
	sm, _ := mem.SwapMemory()

	fmt.Println("------ SwapMemory ------")
	fmt.Printf("Total: %v\n", sm.Total)
	fmt.Printf("Used: %v\n", sm.Used)
	fmt.Printf("Free: %v\n", sm.Free)
	fmt.Printf("UsedPercent: %f%%\n", sm.UsedPercent)
	fmt.Printf("Sin: %v\n", sm.Sin)
	fmt.Printf("Sout: %v\n", sm.Sout)
}

func testDisk() {
	partitions, _ := disk.Partitions(false)

	for _, p := range partitions {
		fmt.Println()
		fmt.Printf("Device: %s\n", p.Device)
		fmt.Printf("Mountpoint: %s\n", p.Mountpoint)
		fmt.Printf("Fstype: %s\n", p.Fstype)
		fmt.Printf("Opts: %s\n", p.Opts)
		fmt.Printf("SerialNumber: %s\n", disk.GetDiskSerialNumber(p.Device))

		fmt.Println("------ UsageStat ------")
		us, _ := disk.Usage(p.Mountpoint)
		fmt.Printf("Path: %s\n", us.Path)
		fmt.Printf("Fstype: %s\n", us.Fstype)
		fmt.Printf("Total: %v\n", humanize.Bytes(us.Total))
		fmt.Printf("Free: %v\n", humanize.Bytes(us.Free))
		fmt.Printf("Used: %v\n", humanize.Bytes(us.Used))
		fmt.Printf("UsedPercent: %f%%\n", us.UsedPercent)
		fmt.Printf("InodesTotal: %v\n", us.InodesTotal)
		fmt.Printf("InodesUsed: %v\n", us.InodesUsed)
		fmt.Printf("InodesFree: %v\n", us.InodesFree)
		fmt.Printf("InodesUsedPercent: %f%%\n", us.InodesUsedPercent)
	}
}

func testInterfaces() {
	interfaces, _ := net.Interfaces()

	for _, i := range interfaces {
		fmt.Println()
		fmt.Printf("MTU: %v\n", i.MTU)
		fmt.Printf("Name: %s\n", i.Name)
		fmt.Printf("HardwareAddr: %s\n", i.HardwareAddr)
		fmt.Printf("Flags: %s\n", i.Flags)
		fmt.Printf("Addrs: %s\n", i.Addrs)
	}
}

func testConnections() {
	conns, _ := net.Connections("all")

	for _, c := range conns {
		fmt.Printf("Fd: %v\n", c.Fd)
		fmt.Printf("Family: %v\n", c.Family)
		fmt.Printf("Type: %v\n", c.Type)
		fmt.Printf("Laddr: %v\n", c.Laddr)
		fmt.Printf("Raddr: %v\n", c.Raddr)
		fmt.Printf("Status: %s\n", c.Status)
		fmt.Printf("Uids: %v\n", c.Uids)
		fmt.Printf("Pid: %v\n", c.Pid)
	}
}

type Process struct {
	Pid        int32   `json: "pid"`
	Name       string  `json: "name"`
	Exe        string  `json: "exe"`
	Cmdline    string  `json: "cmdline"`
	Terminal   string  `json: "terminal"`
	Status     string  `json: "status"`
	Cwd        string  `json: "cwd"`
	Ppid       int32   `json: "ppid"`
	NumThreads int32   `josn: "numThreads"`
	NumFDs     int32   `json: "numfds"`
	Uids       []int32 `json: "uids"`
	Gids       []int32 `json: "gids"`
	CreateTime int64   `json: "createtime"`
}

func ProcessList() ([]Process, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, err
	}
	count := len(pids)
	log.Printf("%v\n", count)

	processes := make([]Process, 0, count)

	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			log.Printf("Pid: %v error\n", pid)
			continue
		}

		ppid, _ := p.Ppid()
		name, _ := p.Name()
		exe, _ := p.Exe()
		cmdline, _ := p.Cmdline()
		ct, _ := p.CreateTime()
		cwd, _ := p.Cwd()
		uids, _ := p.Uids()
		gids, _ := p.Gids()
		terminal, _ := p.Terminal()
		status, _ := p.Status()
		numFds, _ := p.NumFDs()
		numThreads, _ := p.NumThreads()

		proc := Process{
			Pid:        p.Pid,
			Name:       name,
			Exe:        exe,
			Cmdline:    cmdline,
			Terminal:   terminal,
			Status:     status,
			Cwd:        cwd,
			Ppid:       ppid,
			NumThreads: numThreads,
			NumFDs:     numFds,
			Uids:       uids,
			Gids:       gids,
			CreateTime: ct,
		}

		processes = append(processes, proc)
	}
	return processes, nil
}

func main() {
	fmt.Println("*****************************************")
	// testHostInfo()
	// fmt.Println("*****************************************")
	// testUserStat()
	// fmt.Println("*****************************************")
	testCpuTimesStat()
	// fmt.Println("*****************************************")
	// testCpuInfo()
	// fmt.Println("*****************************************")
	// testLoadavg()
	// fmt.Println("*****************************************")
	// testMem()
	// fmt.Println("*****************************************")
	// testDisk()
	// fmt.Println("*****************************************")
	// testInterfaces()
	// fmt.Println("*****************************************")
	// testConnections()
	// fmt.Println("*****************************************")
	// ps, _ := ProcessList()
	// for _, p := range ps {
	// fmt.Printf("Pid: %v Name: %s Exe: %s Cmdline: %s\n", p.Pid, p.Name, p.Exe, p.Cmdline)
	// }
}
