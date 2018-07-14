// copy code from github.com/shirou/gopsutil
// almost the same, but change a lot, especially the returned struct
package sys

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type SysStat struct {
	// /proc/loadavg
	Load1  float64
	Load5  float64
	Load15 float64
	// /proc/meminfo
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
	Active      uint64
	Inactive    uint64
	Buffers     uint64
	Cached      uint64
	Wired       uint64
	Shared      uint64
	// /proc/net/dev
	// sum of all interfaces
	Name        string
	BytesSent   uint64
	BytesRecv   uint64
	PacketsSent uint64
	PacketsRecv uint64
	// /proc/stat
	CPU       string
	User      float64
	System    float64
	Idle      float64
	Nice      float64
	Iowait    float64
	Irq       float64
	Softirq   float64
	Steal     float64
	Guest     float64
	GuestNice float64
	Stolen    float64
}

func GetSysStat() *SysStat {
	stat := new(SysStat)
	stat.cPUTimes()
	stat.netIOCounters()
	stat.virtualMemory()
	stat.loadAvg()
	return stat
}

func (stat *SysStat) cPUTimes() error {
	filename := "/proc/stat"
	var lines = []string{}
	lines, _ = readLinesOffsetN(filename, 0, 1)

	err := stat.parseStatLine(lines[0])
	if err != nil {
		return err
	}
	return nil
}

func (stat *SysStat) parseStatLine(line string) error {
	fields := strings.Fields(line)

	if strings.HasPrefix(fields[0], "cpu") == false {
		return errors.New("not contain cpu")
	}

	cpu := fields[0]
	if cpu == "cpu" {
		cpu = "cpu-total"
	}
	user, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return err
	}
	nice, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return err
	}
	system, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return err
	}
	idle, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return err
	}
	iowait, err := strconv.ParseFloat(fields[5], 64)
	if err != nil {
		return err
	}
	irq, err := strconv.ParseFloat(fields[6], 64)
	if err != nil {
		return err
	}
	softirq, err := strconv.ParseFloat(fields[7], 64)
	if err != nil {
		return err
	}
	stolen, err := strconv.ParseFloat(fields[8], 64)
	if err != nil {
		return err
	}

	cpuTick := float64(100) // TODO: how to get _SC_CLK_TCK ?

	stat.CPU = cpu
	stat.User = float64(user) / cpuTick
	stat.Nice = float64(nice) / cpuTick
	stat.System = float64(system) / cpuTick
	stat.Idle = float64(idle) / cpuTick
	stat.Iowait = float64(iowait) / cpuTick
	stat.Irq = float64(irq) / cpuTick
	stat.Softirq = float64(softirq) / cpuTick
	stat.Stolen = float64(stolen) / cpuTick

	if len(fields) > 9 { // Linux >= 2.6.11
		steal, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			return err
		}
		stat.Steal = float64(steal)
	}
	if len(fields) > 10 { // Linux >= 2.6.24
		guest, err := strconv.ParseFloat(fields[10], 64)
		if err != nil {
			return err
		}
		stat.Guest = float64(guest)
	}
	if len(fields) > 11 { // Linux >= 3.2.0
		guestNice, err := strconv.ParseFloat(fields[11], 64)
		if err != nil {
			return err
		}
		stat.GuestNice = float64(guestNice)
	}

	return nil
}

type netIOCountersStat struct {
	Name        string
	BytesSent   uint64
	BytesRecv   uint64
	PacketsSent uint64
	PacketsRecv uint64
}

func (stat *SysStat) netIOCounters() error {
	filename := "/proc/net/dev"
	lines, err := readLines(filename)
	if err != nil {
		return err
	}

	statlen := len(lines) - 1

	all := make([]netIOCountersStat, 0, statlen)

	for _, line := range lines[2:] {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		interfaceName := strings.TrimSpace(parts[0])
		if interfaceName == "" {
			continue
		}

		fields := strings.Fields(strings.TrimSpace(parts[1]))
		bytesRecv, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return err
		}
		packetsRecv, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return err
		}
		bytesSent, err := strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return err
		}
		packetsSent, err := strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			return err
		}

		nic := netIOCountersStat{
			Name:        interfaceName,
			BytesRecv:   bytesRecv,
			PacketsRecv: packetsRecv,
			BytesSent:   bytesSent,
			PacketsSent: packetsSent}

		all = append(all, nic)
	}

	return stat.getNetIOCountersAll(all)
}

func (stat *SysStat) getNetIOCountersAll(n []netIOCountersStat) error {
	stat.Name = "all-interfaces"
	for _, nic := range n {
		stat.BytesRecv += nic.BytesRecv
		stat.PacketsRecv += nic.PacketsRecv
		stat.BytesSent += nic.BytesSent
		stat.PacketsSent += nic.PacketsSent
	}
	return nil
}

func (stat *SysStat) virtualMemory() error {
	filename := "/proc/meminfo"
	lines, _ := readLines(filename)

	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)

		t, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		switch key {
		case "MemTotal":
			stat.Total = t * 1000
		case "MemFree":
			stat.Free = t * 1000
		case "Buffers":
			stat.Buffers = t * 1000
		case "Cached":
			stat.Cached = t * 1000
		case "Active":
			stat.Active = t * 1000
		case "Inactive":
			stat.Inactive = t * 1000
		}
	}
	stat.Available = stat.Free + stat.Buffers + stat.Cached
	stat.Used = stat.Total - stat.Free
	stat.UsedPercent = float64(stat.Total-stat.Available) / float64(stat.Total) * 100.0

	return nil
}

func (stat *SysStat) loadAvg() error {
	filename := "/proc/loadavg"
	line, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	values := strings.Fields(string(line))

	load1, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return err
	}
	load5, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return err
	}
	load15, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return err
	}

	stat.Load1 = load1
	stat.Load5 = load5
	stat.Load15 = load15

	return nil
}
