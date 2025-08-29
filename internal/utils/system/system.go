package system

import (
	"os"
	"os/user"
	"sprout/internal/utils/logger"
	"time"
)

func CheckSudo() {
    uid := os.Geteuid()
    if uid != 0 {
        logger.Fatal("Permission denied: sudo privileges required")
    }
}

func GetUser() (*user.User, error){
	return user.Lookup(os.Getenv("SUDO_USER"))
}

func RunCommand(command string) (string, error) {
	// cmd := exec.Command("sh", "-c", command)

	// out, err := cmd.CombinedOutput()

	// return string(out), err

	time.Sleep(1 * time.Second)
	return "", nil
}