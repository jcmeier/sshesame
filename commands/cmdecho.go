package commands

import (
	"fmt"
	"strings"
)

type cmdEcho struct{}

func (cmdEcho) execute(context CommandContext) (uint32, error) {
	_, err := fmt.Fprintln(context.stdout, strings.Join(context.args[1:], " "))
	return 0, err
}
