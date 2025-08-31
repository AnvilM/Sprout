package main

import (
	"os/exec"
	"fmt"
)

func main() {
	command := "echo $SUDO_USER"
	cmd := exec.Command("sh", "-c", command)

	out, err := cmd.CombinedOutput()

	fmt.Printf("%s", out)
	if(err != nil){
		fmt.Printf("%s", err)
	}
}