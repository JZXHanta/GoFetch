package main

import (
	"fmt"
	"log"
	"os/user"
	"runtime"
	"strings"

	osinfo "github.com/JZXHanta/OSInfo"
	// "github.com/mohae/joefriday/cpu/cpuinfo"
	"github.com/mackerelio/go-osstat/uptime"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/v3/cpu"
)

func cpuInfo() string {
	var c string
	switch runtime.GOOS {
	case "windows":
		cpuinfo, err := cpu.Info()
		if err != nil {
			log.Fatalf(err.Error())
		}
		//fmt.Print(cpuinfo)
		c = cpuinfo[0].ModelName

	case "linux":
		cpuinfo, err := cpu.Info()
		if err != nil {
			log.Fatalf(err.Error())
		}
		//fmt.Print(cpuinfo)
		c = cpuinfo[0].ModelName

	}
	return fmt.Sprintf("CPU     : %s", c)
}

func upTime() string {
	uptime, _ := uptime.Get()
	str := uptime.String()

	return fmt.Sprintf("Uptime  : %s", str)
}

func osInfo() (string, string) {
	var os string
	var ver string
	switch runtime.GOOS {
	case "linux":
		os = fmt.Sprint("OS      : ", osinfo.LinuxDistro())
		ver = fmt.Sprint("Version : ", osinfo.LinuxVersion())
	case "windows":
		os = fmt.Sprint("OS      : ", osinfo.WindowsOS())
		ver = fmt.Sprint("Version : ", osinfo.WindowsKernel())
	}
	return os, ver
}

func totalMemory() string {
	total := ((memory.TotalMemory() / 1024) / 1024) / 1000
	free := ((memory.FreeMemory() / 1024) / 1024) / 1000
	used := total - free
	return fmt.Sprintf("Ram     : %d / %d GiB", used, total)
}

func userNameAndHostName() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	u := currentUser.Username
	userAndComp := strings.Split(u, "\\")
	userName := userAndComp[1]
	compName := userAndComp[0]

	return fmt.Sprintf("%s @ %s", userName, compName)
}

func goFetch() {
	OS, VER := osInfo()
	USERATHOST := userNameAndHostName()
	SEP := "------------------------"
	MEMUSED := totalMemory()
	CPU := cpuInfo()
	UPTIME := upTime()

	fmt.Println("")
	fmt.Println(USERATHOST) // "User @ Hostname"
	fmt.Println(SEP)        // "-----------------------"
	fmt.Println(OS)         // "OS      :  Microsoft Windows 11 Pro"
	fmt.Println(VER)        // "Version :  22621"
	fmt.Println(MEMUSED)    // "Ram     :  9 / 32 GiB"
	// TODO: (below)
	// Add TOML config file?
	fmt.Println(UPTIME)
	// fmt.Println(PACKAGES)
	// fmt.Println(SHELL)
	// fmt.Println(RESOLUTION)
	// fmt.Println(TERMINAL)
	fmt.Println(CPU) // "CPU     : AMD Ryzen 5 3600 6-Core Processor"
	// fmt.Println(GPU)
	// fmt.Println(DISKUSED)
	fmt.Println("")
	// maj, min, build := windows.RtlGetNtVersionNumbers()
	// fmt.Println(maj, min, build)
}

func main() {
	goFetch()
}
