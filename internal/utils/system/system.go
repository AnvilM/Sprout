package system

import (
	"os"
	"os/exec"
	"os/user"
	"sprout/internal/utils/logger"
	"sprout/internal/config"
)

func CheckSudo() {
    uid := os.Geteuid()
    if uid == 0 {
        logger.Fatal("Dont rus as sudo")
    }
}

func GetUser() (*user.User, error){
	return user.Lookup(os.Getenv("SUDO_USER"))
}

func RunCommand(command config.SetupCommand) (string, error){
	if command.Root {
		return RunCommandAsRoot(command.Command)
	}

	return RunCommandAsUser(command.Command)
}

func RunCommandAsRoot(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	out, err := cmd.CombinedOutput()

	return string(out), err
}

func RunCommandAsUser(command string) (string, error) {
	username := os.Getenv("SUDO_USER")
	cmd := exec.Command("sudo", "-i", "-u", username, "sh", "-c", command)

	out, err := cmd.CombinedOutput()

	return string(out), err
}