package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

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

func GPUInfo() string {
	info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	history, _ := info.Output()
	str := strings.TrimSpace(strings.Replace(string(history), "Name", "", -1))
	return fmt.Sprintf("GPU       : %s", str)
}

// TODO:
// func Disk() {}
