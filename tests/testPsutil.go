package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
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

func testVirtualMemory() {
	// func SwapMemory() (*SwapMemoryStat, error)
	v1, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v1.Total, v1.Free, v1.UsedPercent)
	fmt.Println(v1)

	// func SwapMemory() (*SwapMemoryStat, error)
	v2, _ := mem.SwapMemory()
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v2.Total, v2.Free, v2.UsedPercent)
	fmt.Println(v2)
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
}
