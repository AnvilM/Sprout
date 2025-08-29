package dependencies

import (
	"sprout/internal/config"
	"sprout/internal/installer/pacman"
)

func installPacmanDeps(cfg *config.Config){
	pacmanDeps := collectUniquePacmanDependencies(cfg)

	if len(pacmanDeps) > 0 {
		pacman.Install(pacmanDeps, " Installing build dependencies: ")
	}
}


func removePacmanDeps(cfg *config.Config){
	pacmanDeps := collectUniquePacmanDependencies(cfg)

	if len(pacmanDeps) > 0 {
		pacman.Remove(pacmanDeps, "Removing build dependencies:")
	}
}


func collectUniquePacmanDependencies(cfg *config.Config) []config.PacmanItem {
	unique := make(map[string]config.PacmanItem)

	// collect root pacman names to exclude
	root := make(map[string]struct{})
	for _, p := range cfg.Pacman {
		root[p.Name] = struct{}{}
	}

	// collect from hyprplugins
	for _, hp := range cfg.Hyprplugins {
		if hp.BuildDependencies != nil && hp.BuildDependencies.Pacman != nil {
			for _, dep := range hp.BuildDependencies.Pacman {
				if dep.Name == "" {
					continue
				}
				if _, exists := root[dep.Name]; exists {
					continue // skip root pacman items
				}
				if _, exists := unique[dep.Name]; !exists {
					unique[dep.Name] = dep
				}
			}
		}
	}

	// collect from additional utils
	for _, au := range cfg.AdditionalUtils {
		if au.BuildDependencies != nil && au.BuildDependencies.Pacman != nil {
			for _, dep := range au.BuildDependencies.Pacman {
				if dep.Name == "" {
					continue
				}
				if _, exists := root[dep.Name]; exists {
					continue // skip root pacman items
				}
				if _, exists := unique[dep.Name]; !exists {
					unique[dep.Name] = dep
				}
			}
		}
	}

	// convert map to slice
	result := make([]config.PacmanItem, 0, len(unique))
	for _, v := range unique {
		result = append(result, v)
	}
	return result
}
