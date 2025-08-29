package spinner

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

var GlobalSpinner *spinner.Spinner

func InitSpinner() *spinner.Spinner {
	GlobalSpinner = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	GlobalSpinner.Color("green")
	GlobalSpinner.Prefix = "  "
	return GlobalSpinner
}

func Start() {
	if GlobalSpinner != nil {
		GlobalSpinner.Start()
	}
}

func Stop() {
	if GlobalSpinner != nil {
		GlobalSpinner.Stop()
	}
}

func Print(str string) {
	if(GlobalSpinner == nil){
		return
	}

	if GlobalSpinner.Active() {
		Stop()
		fmt.Print(str + "\n")
		Start()
	} else {
		fmt.Print(str + "\n")
	}
}

func GetSpinner() *spinner.Spinner {
	return GlobalSpinner
}
