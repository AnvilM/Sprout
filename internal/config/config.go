package config

import (
	_ "embed"
	"fmt"
	"sprout/internal/utils/logger"

	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var configYAML []byte

type Config struct {
	Pacman          []PacmanItem      `yaml:"pacman"`
	Hyprplugins     []HyprpluginItem  `yaml:"hyprplugins"`
	AdditionalUtils []AdditionalUtil  `yaml:"additionalUtils"`
	Assets          *Assets           `yaml:"assets,omitempty"`
}

type PacmanItem struct {
	Name          string   `yaml:"name"`
	SetupCommands []string `yaml:"setupCommands,omitempty"`
}

type HyprpluginItem struct {
	Name              string             `yaml:"name"`
	Repository        string             `yaml:"repository"`
	BuildDependencies *BuildDependencies `yaml:"buildDependencies,omitempty"`
}

type AdditionalUtil struct {
	Name              string             `yaml:"name"`
	BuildDependencies *BuildDependencies `yaml:"buildDependencies,omitempty"`
	SetupCommands     []string           `yaml:"setupCommands,omitempty"`
}

type BuildDependencies struct {
	Pacman []PacmanItem `yaml:"pacman,omitempty"`
	// Future expansion: Aur []AurItem `yaml:"aur,omitempty"`
}

type Assets struct {
	Scripts []string `yaml:"scripts,omitempty"`
}

func GetConfig() *Config {
	var cfg Config
	if err := yaml.Unmarshal(configYAML, &cfg); err != nil {
		logger.Fatal("Failed to parse embedded YAML: %v", err)
	}

	validateConfig(&cfg)
	return &cfg
}

func validateConfig(cfg *Config) {
	// Pacman (root level)
	if cfg.Pacman == nil {
		cfg.Pacman = nil
	} else if len(cfg.Pacman) == 0 {
		logger.Fatal("pacman array cannot be empty")
	}
	for i, item := range cfg.Pacman {
		validatePacmanItem(item, fmtIndex("pacman", i))
	}

	// Hyprplugins
	if cfg.Hyprplugins == nil {
		cfg.Hyprplugins = nil
	} else if len(cfg.Hyprplugins) == 0 {
		logger.Fatal("hyprplugins array cannot be empty")
	}
	for i, item := range cfg.Hyprplugins {
		if item.Name == "" {
			logger.Fatal("hyprplugins[%d].name is required", i)
		}
		if item.Repository == "" {
			logger.Fatal("hyprplugins[%d].repository is required", i)
		}
		if item.BuildDependencies != nil {
			validateBuildDependencies(item.BuildDependencies, fmtIndex("hyprplugins", i))
		}
	}

	// AdditionalUtils
	if cfg.AdditionalUtils == nil {
		cfg.AdditionalUtils = nil
	} else if len(cfg.AdditionalUtils) == 0 {
		logger.Fatal("additionalUtils array cannot be empty")
	}
	for i, item := range cfg.AdditionalUtils {
		if item.Name == "" {
			logger.Fatal("additionalUtils[%d].name is required", i)
		}
		if item.SetupCommands != nil && len(item.SetupCommands) == 0 {
			logger.Fatal("additionalUtils[%d].setupCommands cannot be empty", i)
		}
		if item.BuildDependencies != nil {
			validateBuildDependencies(item.BuildDependencies, fmtIndex("additionalUtils", i))
		}
	}

	// Assets
	if cfg.Assets != nil {
		if cfg.Assets.Scripts != nil && len(cfg.Assets.Scripts) == 0 {
			logger.Fatal("assets.scripts cannot be empty if defined")
		}
	}
}

func validateBuildDependencies(bd *BuildDependencies, prefix string) {
	// If defined, it must not be empty
	if len(bd.Pacman) == 0 {
		logger.Fatal("%s.buildDependencies cannot be empty", prefix)
	}

	// Validate pacman items
	if bd.Pacman != nil {
		for i, p := range bd.Pacman {
			validatePacmanItem(p, fmt.Sprintf("%s.buildDependencies.pacman[%d]", prefix, i))
		}
	}
}

func validatePacmanItem(item PacmanItem, prefix string) {
	if item.Name == "" {
		logger.Fatal("%s.name is required", prefix)
	}
	if item.SetupCommands != nil && len(item.SetupCommands) == 0 {
		logger.Fatal("%s.setupCommands cannot be empty", prefix)
	}
}

func fmtIndex(section string, i int) string {
	return fmt.Sprintf("%s[%d]", section, i)
}
