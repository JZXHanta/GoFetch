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
	toml "github.com/pelletier/go-toml"
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
		c = cpuinfo[0].ModelName

	case "linux":
		cpuinfo, err := cpu.Info()
		if err != nil {
			log.Fatalf(err.Error())
		}
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
	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal(err)
	}
	h := fmt.Sprint(strings.TrimSpace(string(out)))
	return fmt.Sprintf("%s @ %s", u, h)
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

func linuxPackages() string {
	os, _ := osInfo()
	var packages string
	switch os {
	case "OS      : Ubuntu":
		packages = aptPackages()
	case "OS      : Pop!_OS":
		PrintLogo(PopOsLogo, OKBLUE)
	case "OS      : Fedora":
		PrintLogo(FedoraLogo, OKBLUE)
	}
	return packages
}

func aptPackages() string {
	cmd := "apt-mark showmanual | wc -l"
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return fmt.Sprint(out)
}

func packageCount() string {
	var packageCount []string
	var str string
	switch runtime.GOOS {
	case "linux":
		pkg := linuxPackages()
		str = fmt.Sprint(pkg) // Sprint will be used properly later...
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

func allInfo() []string {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Visuals
	var USERATHOSTNAME bool = config.Get("VISUALS.USERATHOSTNAME").(bool)
	var SEPARATOR bool = config.Get("VISUALS.SEPARATOR").(bool)
	var CUSTOM_SEPARATOR string = config.Get("VISUALS.CUSTOM_SEPARATOR").(string)
	//var CUSTOM_ART bool = config.Get("VISUALS.CUSTOM_ART").(bool)
	//var CUSTOM_ASCII_ART string = config.Get("VISUALS.CUSTOM_ASCII_ART").(string)

	// Software
	var OS bool = config.Get("SOFTWARE.OS").(bool)
	var VERSION bool = config.Get("SOFTWARE.VERSION").(bool)
	var SHELL bool = config.Get("SOFTWARE.SHELL").(bool)
	var TERMINAL bool = config.Get("SOFTWARE.TERMINAL").(bool)

	// Hardware
	var RAM bool = config.Get("HARDWARE.RAM").(bool)
	var CPU bool = config.Get("HARDWARE.CPU").(bool)
	var GPU bool = config.Get("HARDWARE.GPU").(bool)
	var DISKUSED bool = config.Get("HARDWARE.DISKUSED").(bool)

	// Info
	var UPTIME bool = config.Get("INFO.UPTIME").(bool)
	var PACKAGES bool = config.Get("INFO.PACKAGES").(bool)
	var RESOLUTION bool = config.Get("INFO.RESOLUTION").(bool)

	var array []string

	if USERATHOSTNAME {
		v := userNameAndHostName()
		array = append(array, v)
	}

	var newSep string
	if SEPARATOR && CUSTOM_SEPARATOR != "" {
		sep := CUSTOM_SEPARATOR
		for i := 0; i < 25; i++ {
			newSep += sep
		}
		array = append(array, newSep)
	} else if SEPARATOR {
		sep := "-"
		for i := 0; i < 25; i++ {
			newSep += sep
		}
		array = append(array, newSep)
	} else {
		array = append(array, "")
	}

	if OS {
		v, _ := osInfo()
		array = append(array, v)
	}

	if VERSION {
		_, v := osInfo()
		array = append(array, v)
	}

	if SHELL {
		v := shell()
		array = append(array, v)
	}

	if TERMINAL {
		v := "TERMINAL STILL WIP"
		array = append(array, v)
	}

	if RAM {
		v := totalMemory()
		array = append(array, v)
	}

	if CPU {
		v := cpuInfo()
		array = append(array, v)
	}

	if GPU {
		v := "GPU STILL WIP"
		array = append(array, v)
	}

	if DISKUSED {
		v := "DISK STILL WIP"
		array = append(array, v)
	}

	if UPTIME {
		v := upTime()
		array = append(array, v)
	}

	if PACKAGES {
		v := packageCount()
		array = append(array, v)
	}

	if RESOLUTION {
		v := "RESOLUTION STILL WIP"
		array = append(array, v)
	}

	return array
}

func PrintLogo(logo, color string) {
	arr := allInfo()
	var array = strings.Split(logo, "\n")
	for i := 0; i < len(array); i++ {
		if i < len(arr) {
			fmt.Println(color, array[i], ENDC, arr[i])
		} else {
			fmt.Println(color, array[i], ENDC)
		}
	}
}

func render() {
	o, _ := osInfo()
	switch o {
	case "OS      : Ubuntu":
		PrintLogo(UbuntuLogo, UBUNTUCOLOR)
	case "OS      : Microsoft Windows 11 Pro":
		PrintLogo(WindowsLogo, OKBLUE)
	case "OS      : Pop!_OS":
		PrintLogo(PopOsLogo, OKBLUE)
	case "OS      : Fedora":
		PrintLogo(FedoraLogo, OKBLUE)

	}

}

func goFetch() {
	render()

	// TODO: (below)
	// fmt.Println(RESOLUTION)
	// fmt.Println(TERMINAL)
	// fmt.Println(GPU)
	// fmt.Println(DISKUSED)

}

func main() {
	goFetch()
}
