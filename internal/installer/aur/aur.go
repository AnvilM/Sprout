package aur

import (
	"os"
	"path/filepath"
	"time"

	"anvilarch/internal/utils/logger"
	"anvilarch/internal/utils/system"

	"github.com/briandowns/spinner"
)

var yayBuildDeps = []string{"git", "base-devel", "go"}

func InstallYay(clean bool) {
	steps := []string{
		"Installing build dependencies",
		"Cloning yay repository",
		"Building and installing yay",
	}

	if clean {
		steps = append(steps, "Removing temporary build dependencies")
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	for _, step := range steps {
		s.Suffix = " " + step
		s.Start()

		switch step {
		case "Installing build dependencies":
			if err := system.RunCommand("sudo", append([]string{"pacman", "-S", "--noconfirm"}, yayBuildDeps...)...); err != nil {
				s.Stop()
				logger.Fatal("%s: %v", step, err)
			}

		case "Cloning yay repository":
			if err := system.RunCommand("git", "clone", "https://aur.archlinux.org/yay.git"); err != nil {
				s.Stop()
				logger.Fatal("%s: %v", step, err)
			}

		case "Building and installing yay":
			yayPath := filepath.Join(".", "yay")
			if err := os.Chdir(yayPath); err != nil {
				s.Stop()
				logger.Fatal("Failed to change directory to %s: %v", yayPath, err)
			}
			if err := system.RunCommand("makepkg", "-si", "--noconfirm", "--needed"); err != nil {
				s.Stop()
				logger.Fatal("%s: %v", step, err)
			}

		case "Removing temporary build dependencies":
			if err := system.RunCommand("sudo", append([]string{"pacman", "-Rs", "--noconfirm"}, yayBuildDeps...)...); err != nil {
				s.Stop()
				logger.Fatal("%s: %v", step, err)
			}
		}

		s.Stop()
	}

	logger.Log(logger.SUCCESS, "Yay has been successfully installed.")

}
