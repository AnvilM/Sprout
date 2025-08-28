package hypr_plugins

import (
	_ "embed"
	"fmt"
	"foundryinstaller/internal/utils/logger"
	"foundryinstaller/internal/utils/system"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var pluginsYAML []byte

type Plugin struct {
    Name          string   `yaml:"name"`
    Repository    string   `yaml:"repository"`
}

var buildDeps = []string{"cpio", "cmake", "git", "meson", "gcc"}

func InstallPlugins(clean bool) {

	var plugins []Plugin
    if err := yaml.Unmarshal(pluginsYAML, &plugins); err != nil {
        log.Fatal(err)
    }

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	for _, pkg := range buildDeps {
		s.Suffix = fmt.Sprintf(" Installing build dependencies: %s", pkg)
		if err := system.RunCommand("pacman -S" + pkg); err != nil{
			s.Stop()
			logger.Fatal("%s", err)
		}
	}

    for _, plugin := range plugins {

        s.Suffix = fmt.Sprintf(" Installing plugins: %s", plugin.Name)

		if err := system.RunCommand(`sudo -u "$SUDO_USER" -i bash -c "hyprpm add ` + plugin.Repository +`"`); err != nil {
			s.Stop()
			logger.Fatal("%s", err)
		}

		if err := system.RunCommand(`sudo -u "$SUDO_USER" -i bash -c "hyprpm enable ` + plugin.Name +`"`); err != nil {
			s.Stop()
			logger.Fatal("%s", err)
		}
    }

	if(clean){
		for _, pkg := range buildDeps {
			s.Suffix = fmt.Sprintf(" Removing build dependencies: %s", pkg)
			if err := system.RunCommand("pacman -Rns" + pkg); err != nil{
				s.Stop()
				logger.Fatal("%s", err)
			}
		}
	}

	s.Stop()

	logger.Log(logger.SUCCESS, "All hyprland plugins successfully installed!")
}


