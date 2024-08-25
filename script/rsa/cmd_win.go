//go:build win
// +build win

package main

import "os/exec"

func Execute(shell string) *exec.Cmd {
	return exec.Command("cmd", "/C", shell)
}
