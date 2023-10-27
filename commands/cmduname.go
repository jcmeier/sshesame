package commands

import (
	"fmt"
)

type cmdUname struct{}

const uname = "Linux shelx 5.15.0-87-generic #97-Ubuntu SMP Mon Oct 2 21:09:21 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux"

func (cmdUname) execute(context CommandContext) (uint32, error) {
	result := uname
	_, err := fmt.Fprintln(context.stdout, result)
	return 0, err
}
