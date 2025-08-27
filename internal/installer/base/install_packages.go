package base

import (
	"anvilarch/internal/utils/logger"
	"anvilarch/internal/utils/system"
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var packagesYAML []byte

type Package struct {
    Name          string   `yaml:"name"`
    SetupCommands []string `yaml:"setupCommands,omitempty"`
}

func InstallPackages() {

	var pkgs []Package
    if err := yaml.Unmarshal(packagesYAML, &pkgs); err != nil {
        log.Fatal(err)
    }

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

    for _, p := range pkgs {

        s.Suffix = fmt.Sprintf(" Installing packages: %s", p.Name)

		if err := system.RunCommand("pacman -S " + p.Name); err != nil {
			s.Stop()
			logger.Fatal("%s", err)
		}

        for _, cmd := range p.SetupCommands {
			if err := system.RunCommand(cmd); err != nil {
				s.Stop()
				logger.Fatal("%s", err)
			}
        }
    }

	s.Stop()

	logger.Log(logger.SUCCESS, "All packages successfully installed!")
}
