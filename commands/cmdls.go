package commands

import (
	"fmt"
	"strings"
)

type cmdLs struct{}

var userfiles = []string{"Hello", "World", "Golang", "Example"}

func (cmdLs) execute(context CommandContext) (uint32, error) {
	separator := " "
	result := strings.Join(userfiles, separator) + " ARGS: "
	result += strings.Join(context.args, " ")
	_, err := fmt.Fprintln(context.stdout, result)
	return 0, err
}
