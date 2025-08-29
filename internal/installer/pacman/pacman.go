package pacman

import (
	"sprout/internal/config"
	"sprout/internal/utils/logger"
	"sprout/internal/utils/system"
	"sprout/internal/utils/tui/spinner"
)


func Install(pacman []config.PacmanItem, label string){

    for _, pacmanItem := range pacman {

        spinner.GetSpinner().Suffix = label + pacmanItem.Name

		if _, err := system.RunCommand("pacman -S " + pacmanItem.Name); err != nil {
			logger.Fatal("%s", err)
		}

        for _, cmd := range pacmanItem.SetupCommands {
			if _, err := system.RunCommand(cmd); err != nil {
				logger.Fatal("%s", err)
			}
        }
    }
}


func Remove(pacman []config.PacmanItem, label string){

    for _, pacmanItem := range pacman {
        spinner.GetSpinner().Suffix = label + pacmanItem.Name

		if _, err := system.RunCommand("pacman -Rns " + pacmanItem.Name); err != nil {
			logger.Fatal("%s", err)
		}
    }
}