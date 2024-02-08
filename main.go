package main

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"

	osinfo "github.com/JZXHanta/OSInfo"
	"github.com/pbnjay/memory"
)

func moreInfo() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("OS      : ", osinfo.LinuxDistro())
		fmt.Println("Version : ", osinfo.LinuxVersion())
	case "windows":
		fmt.Println("OS      : ", osinfo.WindowsOS())
		fmt.Println("Version : ", osinfo.WindowsKernel())
	}
}

func totalMemory() string {
	total := ((memory.TotalMemory() / 1024) / 1024) / 1000
	free := ((memory.FreeMemory() / 1024) / 1024) / 1000
	used := total - free
	t := int(total)
	str := fmt.Sprintf("Ram     :  %d / %d GiB", used, t)
	return str
}

func main() {
	// fmt.Println(getPlatform())
	// fmt.Println(get_arch())
	// readLinuxReleaseInfo()
	moreInfo()
	fmt.Println(totalMemory())
	computerName, _ := windows.ComputerName()
	fmt.Println(computerName)
	maj, min, build := windows.RtlGetNtVersionNumbers()
	fmt.Println(maj, min, build)
}
