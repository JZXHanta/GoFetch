package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	osinfo "github.com/JZXHanta/OSInfo"
)

func ChocoPackages() (string, error) {
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

func WingetPackages() (string, error) {
	out, err := exec.Command("powershell", "-NoProfile", "-NonInteractive", "(winget list --scope user).Count").Output()
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

func LinuxPackages() string {
	os, _ := OsInfo()
	var packages string
	switch os {
	case "OS      : Ubuntu":
		packages = AptPackages()
	case "OS      : Pop!_OS":
		PrintLogo(PopOsLogo, OKBLUE)
	case "OS      : Fedora":
		packages = DnfPackages()
	}
	return packages
}

func AptPackages() string {
	//cmd := "'apt-mark showmanual'"

	out, err := exec.Command("apt-mark", "showmanual").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	str := strings.Split(string(out), "\n")
	count := len(str)

	return fmt.Sprint(count)
}

func DnfPackages() string {
	//cmd := "'apt-mark showmanual'"

	out, err := exec.Command("dnf", "list", "--installed").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	str := strings.Split(string(out), "\n")
	count := len(str)

	return fmt.Sprint(count)
}

func OsInfo() (string, string) {
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

func PackageCount() string {
	var PackageCount []string
	var str string
	switch runtime.GOOS {
	case "linux":
		pkg := LinuxPackages()
		str = fmt.Sprint(pkg)
	case "windows":
		chocoCount, chocoErr := ChocoPackages()
		if chocoErr != nil {
			log.Fatalf(chocoErr.Error())
		} else {
			PackageCount = append(PackageCount, chocoCount)
		}

		wingetCount, wingetErr := WingetPackages()
		if wingetErr != nil {
			log.Fatalf(wingetErr.Error())
		} else {
			PackageCount = append(PackageCount, wingetCount)
		}

		if len(PackageCount) > 0 {
			str = strings.Join(PackageCount, ", ")
		} else {
			str = PackageCount[0]
		}

	}

	return fmt.Sprintf("Packages: %s", str)
}

func Shell() string {
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

// TODO:
// func Terminal() {}
