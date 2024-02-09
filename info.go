package main

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/fstanis/screenresolution"
	"github.com/mackerelio/go-osstat/uptime"
)

func UpTime() string {
	uptime, _ := uptime.Get()
	str := uptime.String()

	return fmt.Sprintf("Uptime    : %s", str)
}

func UserNameAndHostName() string {
	var str string
	switch runtime.GOOS {
	case "linux":
		str = UserNameLinux()
	case "windows":
		str = UserNameWindows()
	}
	return str
}

func UserNameWindows() string {
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

func UserNameLinux() string {
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

func Resolution() string {
	res := screenresolution.GetPrimary()
	return fmt.Sprintf("Resolution: %dx%d", res.Width, res.Height)
}
