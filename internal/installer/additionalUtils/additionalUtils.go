package additionalutils

import (
	"sprout/internal/config"
	"sprout/internal/utils/logger"
	"sprout/internal/utils/system"
	"sprout/internal/utils/tui/spinner"
)


func Install(additionalutils []config.AdditionalUtil){
	spinner.Print(logger.FormatInfo(" Installing additional utils!"))

    for _, additionalutil := range additionalutils {
		spinner.GetSpinner().Suffix = " Installing additional utils: " + additionalutil.Name

        for _, cmd := range additionalutil.SetupCommands {
			if _, err := system.RunCommand(cmd); err != nil {
				logger.Fatal("%s", err)
			}
        }
    }

	spinner.Print(logger.FormatSuccess(" All additional utils successfully installed!\n"))
}