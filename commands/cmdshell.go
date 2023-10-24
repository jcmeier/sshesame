package commands

import (
	"fmt"
	"strconv"
	"strings"
)

type cmdShell struct{}

func (cmdShell) execute(context CommandContext) (uint32, error) {
	var prompt string
	if context.pty {
		switch context.user {
		case "root":
			prompt = "# "
		default:
			prompt = "$ "
		}
	}
	var lastStatus uint32
	var line string
	var err error
	for {
		_, err = fmt.Fprint(context.stdout, prompt)
		if err != nil {
			return lastStatus, err
		}
		line, err = context.stdin.ReadLine()
		if err != nil {
			return lastStatus, err
		}
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		if args[0] == "exit" {
			var err error
			var status uint64 = uint64(lastStatus)
			if len(args) > 1 {
				status, err = strconv.ParseUint(args[1], 10, 32)
				if err != nil {
					status = 255
				}
			}
			return uint32(status), nil
		}
		newContext := context
		newContext.args = args
		if lastStatus, err = ExecuteProgram(newContext); err != nil {
			return lastStatus, err
		}
	}
}
