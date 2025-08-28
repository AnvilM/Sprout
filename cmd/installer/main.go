package main

import (
	"foundryinstaller/internal/installer/base"
	"foundryinstaller/internal/installer/hypr_plugins"
	"foundryinstaller/internal/utils/system"
	"foundryinstaller/internal/utils/tui"
)

func main() {
	system.CheckSudo()


	clean := tui.SelectBool("Remove build dependencies?")
	
	
	base.InstallPackages()

	base.ExportConfigs()

	hypr_plugins.InstallPlugins(clean)

	// if(installYay){
	// 	aur.InstallYay(clean)
	// }

}