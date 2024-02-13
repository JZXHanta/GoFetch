package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/v3/cpu"
)

func CpuInfo() string {
	var c string
	switch runtime.GOOS {
	case "windows":
		cpuinfo, err := cpu.Info()
		if err != nil {
			log.Fatalf(err.Error())
		}
		c = cpuinfo[0].ModelName

	case "linux":
		cpuinfo, err := cpu.Info()
		if err != nil {
			log.Fatalf(err.Error())
		}
		c = cpuinfo[0].ModelName

	}
	return fmt.Sprintf("CPU       : %s", c)
}

func TotalMemory() string {
	total := ((memory.TotalMemory() / 1024) / 1024) / 1000
	free := ((memory.FreeMemory() / 1024) / 1024) / 1000
	used := total - free
	return fmt.Sprintf("Ram       : %d / %d GiB", used, total)
}

func GPUInfoWindows() string {
	info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	// info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	history, _ := info.Output()
	str := strings.TrimSpace(strings.Replace(string(history), "Name", "", -1))
	return fmt.Sprintf("GPU       : %s", str)
}

func GPUInfo() string {
	var str string
	switch runtime.GOOS {
	case "linux":
		str = "WIP"
	case "windows":
		str = GPUInfoWindows()
	}
	return str
}

func DiskInfoWindows() string {
	cmd, err := exec.Command("powershell", "-NoProfile", "-NonInteractive", "wmic diskdrive get Size").Output()
	if err != nil {
		log.Fatalf(err.Error(), "Disk info not working...")
	}
	str := string(cmd)
	list := strings.Split(str, "\n")
	var arr []string
	for i := 0; i < len(list); i++ {
		j := strings.TrimSpace(list[i])
		arr = append(arr, j)
	}
	var sum int64
	for i := 1; i < len(arr); i++ {
		// Below works but if I don't ignore the error it fails for some reason
		j, _ := strconv.ParseInt(arr[i], 10, 0)
		sum += j
	}
	total := ((sum / 1024) / 1024) / 1024
	r := GBtoString(total)
	return fmt.Sprintf("DISK      : %s total, %d drives", r, (len(arr) - 3))
}

func DiskInfoLinux() string {
	cmd := "df --total | grep total"
	out, err := exec.Command(cmd).Output()
	s := string(out)
	if err != nil {
		log.Fatalf(err.Error())
	}
	o := strings.TrimSpace(s)
	list := strings.Split(o, "\n")
	used, _ := strconv.ParseInt(list[2], 10, 0)
	avail, _ := strconv.ParseInt(list[3], 10, 0)
	total := bytesToGB(used + avail)

	return fmt.Sprintf("DISK      : %s", GBtoString(total))
}

func bytesToGB(bytes int64) int64 {
	return ((bytes / 1024) / 1024) / 1024
}

func GBtoString(gb int64) string {
	var str string
	if gb > 1000 {
		str = fmt.Sprintf("%d TiB", (gb / 1024))
	} else {
		str = fmt.Sprintf("%d GiB", gb)
	}
	return str
}

func DiskInfo() string {
	var str string
	switch runtime.GOOS {
	case "windows":
		str = DiskInfoWindows()
	case "linux":
		str = DiskInfoLinux()
	}
	return str
}

// TODO:
// func Disk() {}
