package installer

import (
	"sprout/internal/config"
	"sprout/internal/dependencies"
	additionalutils "sprout/internal/installer/additionalUtils"
	"sprout/internal/installer/hyprplugins"
	"sprout/internal/installer/pacman"
	"sprout/internal/utils/logger"
	"sprout/internal/utils/tui"
	"sprout/internal/utils/tui/spinner"
)


func Install(){
	config := config.GetConfig()

	clean := tui.SelectBool("Remove build dependencies?")

	spinner.InitSpinner().Suffix = " Installing";
	spinner.Start()



	// Install dependencies
	dependencies.Install(config)



	// Install pacman packages
	if config.Pacman != nil {
		spinner.Print(logger.FormatInfo(" Installing packages"))
		pacman.Install(config.Pacman, " Installing packages: ")
		spinner.Print(logger.FormatSuccess(" All pacman packages successfully installed!\n"))
	}



	// Install hyprland plugins
	if config.Hyprplugins != nil {
		hyprplugins.Install(config.Hyprplugins)
	}



	// Install additional utilities
	if config.AdditionalUtils != nil {
		additionalutils.Install(config.AdditionalUtils)
	}



	// Remove dependencies
	if clean {
		dependencies.Remove(config)
	}


	
	spinner.Stop()
}




