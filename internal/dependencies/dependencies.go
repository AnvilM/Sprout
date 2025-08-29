package dependencies

import (
	"sprout/internal/config"
	"sprout/internal/utils/logger"
	"sprout/internal/utils/tui/spinner"
)

func Install(cfg *config.Config){
	spinner.Print(logger.FormatInfo(" Installing build dependencies"))

	installPacmanDeps(cfg)

	spinner.Print(logger.FormatSuccess(" All build dependencies successfully installed!\n"))
}

func Remove(cfg *config.Config){
	spinner.Print(logger.FormatInfo(" Removing build dependencies"))
	removePacmanDeps(cfg)
	spinner.Print(logger.FormatSuccess(" All build dependencies successfully removed!\n"))
}