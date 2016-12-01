package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
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

	percent, _ := cpu.Percent(1*time.Second, true)
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
	d, _ := disk.Usage("/")
	fmt.Printf("HD : %vGB Free : %vGB Usage : %f%%\n", d.Total/1024^3, d.Free/1024^3, d.UsedPercent)
}

func main() {
	fmt.Println("*****************************************")
	testHostInfo()
	fmt.Println("*****************************************")
	testUserStat()
	fmt.Println("*****************************************")
	testCpuTimesStat()
	fmt.Println("*****************************************")
	testCpuInfo()
	fmt.Println("*****************************************")
	testLoadavg()
	fmt.Println("*****************************************")
	testMem()
	fmt.Println("*****************************************")
	testDisk()
}
