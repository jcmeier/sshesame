package commands

import (
	"fmt"
)

type cmdCd struct{}

func (cmdCd) execute(context CommandContext) (uint32, error) {
	result := ""
	_, err := fmt.Fprint(context.stdout, result)
	return 0, err
}
