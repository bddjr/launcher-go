//go:build windows

package main

import (
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command(`.\launch.bat`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}
