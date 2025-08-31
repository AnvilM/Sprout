package main

import (
	"sprout/internal/installer"
	"sprout/internal/utils/system"
	"os/exec"
)

func main() {
	command := "echo $SUDO_USER"
	cmd := exec.Command("sh", "-c", command)

	installer.Install()
}