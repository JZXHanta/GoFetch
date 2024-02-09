package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	osinfo "github.com/JZXHanta/OSInfo"
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

// TODO: Does not work properly in linux, fails with index error on line 76
func userNameAndHostName() string {
	var str string
	switch runtime.GOOS {
	case "linux":
		str = userNameLinux()
	case "windows":
		str = userNameWindows()
	}
	return str
}

func userNameWindows() string {
	var userName string
	var compName string
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	u := currentUser.Username
	userAndComp := strings.Split(u, "\\")
	userName = userAndComp[1]
	compName = userAndComp[0]

	return fmt.Sprintf("%s @ %s", userName, compName)
}

func userNameLinux() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	u := currentUser.Username
	return fmt.Sprintf("%s @ hostname", u)
}

func chocoPackages() (string, error) {
	out, err := exec.Command("powershell", "-NoProfile", "-NonInteractive", "(ls C:\\ProgramData\\chocolatey\\lib).Count").Output()
	if err != nil {
		log.Fatal(err)
	}
	c := fmt.Sprint(string(out))
	count, err := strconv.ParseInt(strings.TrimSpace(c), 10, 10)
	if err != nil {
		log.Fatalf(err.Error())
		return "", err
	}

	return fmt.Sprintf("(Choco: %d)", count), nil
}

func wingetPackages() (string, error) {
	out, err := exec.Command("powershell", "-NoProfile", "-NonInteractive", "(winget list).Count").Output()
	if err != nil {
		log.Fatal(err)
	}
	c := fmt.Sprint(string(out))
	count, err := strconv.ParseInt(strings.TrimSpace(c), 10, 10)
	if err != nil {
		log.Fatalf(err.Error())
		return "", err
	}

	return fmt.Sprintf("(Winget: %d)", count), nil
}

func packageCount() string {
	var packageCount []string
	var str string
	switch runtime.GOOS {
	case "linux":
		str = fmt.Sprint("WIP")
	case "windows":
		chocoCount, chocoErr := chocoPackages()
		if chocoErr != nil {
			packageCount = append(packageCount, chocoCount)
		}
		wingetCount, wingetErr := wingetPackages()
		if wingetErr != nil {
			packageCount = append(packageCount, wingetCount)
		}
		packageCount = append(packageCount, chocoCount)
		if len(packageCount) > 0 {
			str = strings.Join(packageCount, ", ")
		} else {
			str = packageCount[0]
		}

	}

	return fmt.Sprintf("Packages: %s", str)
}

func shell() string {
	var shell string
	switch runtime.GOOS {
	case "windows":
		shell = "pwsh"
	case "linux":
		s := os.Getenv("SHELL")
		shell = strings.Replace(s, "/bin/", "", -1)
	}
	return fmt.Sprintf("Shell   : %s", shell)
}

func goFetch() {
	OS, VER := osInfo()
	USERATHOST := userNameAndHostName()
	SEP := "------------------------"
	MEMUSED := totalMemory()
	CPU := cpuInfo()
	UPTIME := upTime()
	PACKAGES := packageCount()
	SHELL := shell()

	fmt.Println("")
	fmt.Println(USERATHOST) // "User @ Hostname"
	fmt.Println(SEP)        // "-----------------------"
	fmt.Println(OS)         // "OS      :  Microsoft Windows 11 Pro"
	fmt.Println(VER)        // "Version :  22621"
	fmt.Println(MEMUSED)    // "Ram     :  9 / 32 GiB"
	fmt.Println(UPTIME)     // "Uptime  :  43h45m45.64s"
	fmt.Println(CPU)        // "CPU     :  AMD Ryzen 5 3600 6-Core Processor"
	fmt.Println(PACKAGES)   // "Packages: (Choco: 19)"
	fmt.Println(SHELL)      // "Shell   : pwsh"
	fmt.Println("")
	// TODO: (below)
	// Add TOML config file?
	// fmt.Println(RESOLUTION)
	// fmt.Println(TERMINAL)
	// fmt.Println(GPU)
	// fmt.Println(DISKUSED)
	// maj, min, build := windows.RtlGetNtVersionNumbers()
	// fmt.Println(maj, min, build)
}

func main() {
	goFetch()
}
