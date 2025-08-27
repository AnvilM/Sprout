package base

import (
	"anvilarch/internal/utils/logger"
	"anvilarch/internal/utils/system"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var user = os.Getenv("SUDO_USER")

func Setup(){
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Running daemons"

	s.Start()

	if err := runDaemons(); err != nil {
		s.Stop()
		logger.Fatal("%s", err)
	}
	logger.Log(logger.SUCCESS, "All daemons have successfully run!")

	s.Suffix = " Setup fish"
	if err := setupFish(); err != nil {
		s.Stop()
		logger.Fatal("%s", err)
	}
	logger.Log(logger.SUCCESS, "Fish shell has been successfully set up!")
}

func runDaemons() error{
	

	if err := system.RunCommand("systemctl", "enable", "NetworkManager"); err != nil {
		return err
	}

	if err := system.RunCommand("sudo", "-u", user, "systemctl", "--user", "enable", "pipewire", "pipewire-pulse", "wireplumber"); err != nil {
		return err
	}

	if err := system.RunCommand("sudo", "-u", user, "systemctl", "--user", "enable", "xdg-desktop-portal-hyprland.service"); err != nil {
		return err
	}

	if err := system.RunCommand("sudo", "-u", user, "swww-daemon"); err != nil {
		return err
	}

	

	return  nil;
}

func setupFish() error{
	if err := system.RunCommand("sudo", "-u", user, "chsh", "-s", "/usr/bin/fish"); err != nil {
		return err
	}

	return nil;
}