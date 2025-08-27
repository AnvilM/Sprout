package main

import (
	"anvilarch/internal/installer/aur"
	"anvilarch/internal/installer/base"
	"anvilarch/internal/installer/hypr_plugins"
	"anvilarch/internal/utils/system"
	"anvilarch/internal/utils/tui"
)

func main() {

	system.CheckRoot()

	clean := tui.SelectBool("Remove build dependencies?")
	installYay := tui.SelectBool("Install yay?")
	
	
	base.InstallPackages()

	base.ExportConfigs()

	hypr_plugins.InstallPlugins(clean)

	if(installYay){
		aur.InstallYay(clean)
	}

	base.Setup()

}