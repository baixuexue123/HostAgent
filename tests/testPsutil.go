package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

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

func main() {
	n, _ := host.Info()
	fmt.Printf("Platform : %v %v\n", n.Platform, n.PlatformVersion)
	fmt.Printf("KernelVersion : %v\n", n.KernelVersion)
	fmt.Printf("Hostname : %v\n", n.Hostname)
	fmt.Printf("BootTime : %v\n", n.BootTime)
}
