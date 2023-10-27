package commands

import (
	"fmt"
	"strings"
)

type cmdLs struct{}

var userFiles = []string{"pictures", "private", "work", "data"}
var rootFiles = []string{"private", "config"}

func (cmdLs) execute(context CommandContext) (uint32, error) {
	separator := " "
	result := ""

	if context.user == "root" {
		result = strings.Join(rootFiles, separator)

	} else {
		result = strings.Join(userFiles, separator)
	}

	_, err := fmt.Fprintln(context.stdout, result)
	return 0, err
}
