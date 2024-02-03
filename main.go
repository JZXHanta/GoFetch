package main

import (
	"fmt"

	// "strings"
	"runtime"

	osinfo "github.com/JZXHanta/OSInfo"
)

// TODO: on powershell systeminfo.exe | findstr /C:"OS" returns os info
// TODO: on bash lsb_release -a returns os info
// TODO: switch runtime.GOOS {
//	case "windows": ..
//  case "linux": ..
//  ...
//}

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

// func getPlatform() string {
// 	const os string = runtime.GOOS
// 	return os
// }

// func get_arch() string {
// 	const arch string = runtime.GOARCH
// 	return arch
// }

func main() {
	// fmt.Println(getPlatform())
	// fmt.Println(get_arch())
	// readLinuxReleaseInfo()
	moreInfo()
}
