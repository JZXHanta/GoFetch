package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	// "strings"
)

func readLinuxReleaseInfo() {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(lines)

}

func getPlatform() string {
	const os string = runtime.GOOS
	return os
}

func get_arch() string {
	const arch string = runtime.GOARCH
	return arch
}

func main() {
	fmt.Println(getPlatform())
	fmt.Println(get_arch())
	readLinuxReleaseInfo()
}
