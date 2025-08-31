package hyprplugins

import (
	"sprout/internal/config"
	"sprout/internal/utils/logger"
	"sprout/internal/utils/system"
	"sprout/internal/utils/tui/spinner"
)

func Install(hyprplugins []config.HyprpluginItem){
	spinner.Print(logger.FormatInfo(" Installing hyprland plugins"))

	for _, plugin := range hyprplugins {
        spinner.GetSpinner().Suffix = " Installing hyprland plugins: " + plugin.Name

		if _, err := system.RunCommandAsUser("hyprpm add " + plugin.Repository); err != nil {
			logger.Fatal("%s", err)
		}

		if _, err := system.RunCommandAsUser("hyprpm enable " + plugin.Name); err != nil {
			logger.Fatal("%s", err)
		}
    }

	spinner.Print(logger.FormatSuccess(" All hyprland plugins successfully installed!\n"))
}