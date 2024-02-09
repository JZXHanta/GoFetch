package main

import (
	"fmt"
	"strings"

	toml "github.com/pelletier/go-toml"
)

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
		v := UserNameAndHostName()
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
		v, _ := OsInfo()
		array = append(array, v)
	}

	if VERSION {
		_, v := OsInfo()
		array = append(array, v)
	}

	if SHELL {
		v := Shell()
		array = append(array, v)
	}

	if TERMINAL {
		v := "TERMINAL STILL WIP"
		array = append(array, v)
	}

	if RAM {
		v := TotalMemory()
		array = append(array, v)
	}

	if CPU {
		v := CpuInfo()
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
		v := UpTime()
		array = append(array, v)
	}

	if PACKAGES {
		v := PackageCount()
		array = append(array, v)
	}

	if RESOLUTION {
		v := Resolution()
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
	o, _ := OsInfo()
	switch o {
	case "OS        : Ubuntu":
		PrintLogo(UbuntuLogo, UBUNTUCOLOR)
	case "OS        : Microsoft Windows 11 Pro":
		PrintLogo(WindowsLogo, OKBLUE)
	case "OS        : Pop!_OS":
		PrintLogo(PopOsLogo, OKBLUE)
	case "OS        : Fedora":
		PrintLogo(FedoraLogo, OKBLUE)

	}

}

func goFetch() {
	render()
}

func main() {
	goFetch()
}
