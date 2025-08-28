package system

import (
	"foundryinstaller/internal/utils/logger"
	"os"
	"os/exec"
)

func CheckSudo() {
    uid := os.Geteuid()
    if uid != 0 {
        logger.Fatal("Permission denied: sudo privileges required")
    }
}

func RunCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)

	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}