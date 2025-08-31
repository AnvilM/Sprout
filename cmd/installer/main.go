package main

import (
	"sprout/internal/installer"
	"sprout/internal/utils/system"
)

func main() {
	 
	system.CheckSudo()

	installer.Install()
}