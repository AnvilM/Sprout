package base

import (
	"fmt"
	"time"

	"anvilarch/internal/utils/logger"
	"anvilarch/internal/utils/system"

	"github.com/briandowns/spinner"
)

func InstallPackages() {
	allPkgs := []string{
		"cliphist", "fish", "grim", "hyprland", "kitty", "less",
		"networkmanager", "pipewire-pulse", "rofi", "slurp", "swww",
		"xdg-desktop-portal-hyprland", "waybar", "wireplumber",
		"noto-fonts-emoji", "ttf-jetbrains-mono", "ttf-nerd-fonts-symbols-mono",
		"ttf-roboto",
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	for _, pkg := range allPkgs {
		s.Suffix = fmt.Sprintf(" Installing packages: %s", pkg)
		
		if err := system.RunCommand("pacman", "-S", "--noconfirm", pkg); err != nil {
			s.Stop()
			logger.Fatal("%s: %v", pkg, err)
		}
		
	}
	s.Stop()

	logger.Log(logger.SUCCESS, "All packages successfully installed!")
	
}
