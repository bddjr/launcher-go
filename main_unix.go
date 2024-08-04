//go:build unix

package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println(syscall.Exec("/bin/bash", []string{"-c", "firefox"}, os.Environ()))
}
