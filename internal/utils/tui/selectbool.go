package tui

import (
	"sprout/internal/utils/logger"

	"github.com/manifoldco/promptui"
)

var options = []struct {
	Label string
	Value bool
}{
	{"Yes", true},
	{"No", false},
}

func SelectBool(label string) bool {
	items := []string{options[0].Label, options[1].Label}
	p := promptui.Select{
		Label: label,
		Items: items,
		HideSelected: true,
	}
	i, _, err := p.Run()
	if err != nil {
		logger.Fatal("%v\n", err)
		return false
	}
	return options[i].Value
}