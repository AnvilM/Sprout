package main

import (
	"anvilarch/internal/installer/base"
	"anvilarch/internal/installer/hypr_plugins"
	"anvilarch/internal/utils/system"
	"anvilarch/internal/utils/tui"
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