package hypr_plugins

import (
	"anvilarch/internal/utils/logger"
	"anvilarch/internal/utils/system"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

var buildDeps = []string{"cpio", "cmake", "git", "meson", "gcc"}

var plugins = map[string]string{
	"Hyprspace": "https://github.com/KZDKM/Hyprspace",
}

func InstallPlugins(clean bool) {

	user := os.Getenv("SUDO_USER")
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	for _, pkg := range buildDeps {
		s.Suffix = fmt.Sprintf(" Installing build dependencies: %s", pkg)
		if err := system.RunCommand("pacman", "-S", pkg); err != nil{
			s.Stop()
			logger.Fatal("%s: %v", pkg, err)
		}
	}


	for name, link := range plugins {
		s.Suffix = fmt.Sprintf(" Installing plugins: %s", name)

		if err := system.RunCommand("sudo", "-u", user, "hyprpm", "add", link); err != nil {
			s.Stop()
			logger.Fatal("%s: %v", link, err)
		}

		if err := system.RunCommand("sudo", "-u", user, "hyprpm", "enable", name); err != nil {
			s.Stop()
			logger.Fatal("%s: %v", name, err)
		}
	}

	if(clean){
		for _, pkg := range buildDeps {
			s.Suffix = fmt.Sprintf(" Removing build dependencies: %s", pkg)
			if err := system.RunCommand("pacman", "-Rns", pkg); err != nil{
				s.Stop()
				logger.Fatal("%s: %v", pkg, err)
			}
		}
	}

	
	s.Stop()

	logger.Log(logger.SUCCESS, "All hyprland plugins successfully installed!")
	
	
}
