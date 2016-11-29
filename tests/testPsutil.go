package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
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

func testCpu() {
	count, _ := cpu.Counts(true)
	fmt.Printf("Cpu Counts: %v\n", count)

	info, _ := cpu.Info()
	fmt.Println(info)
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
}
