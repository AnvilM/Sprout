package system

import (
	"anvilarch/internal/utils/logger"
	"os"
	"time"
)

func CheckRoot() {
	if os.Geteuid() != 0 {
		logger.Fatal("Permission denied: sudo privileges required.")
		os.Exit(1)
	}
}

func RunCommand(name string, args ...string) error {
	// logger.Log(logger.INFO, "run", name, args);
	// cmd := exec.Command(name, args...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// return cmd.Run()
	time.Sleep(1 * time.Second)
	return nil
}