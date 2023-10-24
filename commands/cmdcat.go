package commands

import "fmt"

type cmdCat struct{}

func (cmdCat) execute(context CommandContext) (uint32, error) {
	if len(context.args) > 1 {
		for _, file := range context.args[1:] {
			if _, err := fmt.Fprintf(context.stderr, "%v: %v: No such file or directory\n", context.args[0], file); err != nil {
				return 0, err
			}
		}
		return 1, nil
	}
	var line string
	var err error
	for err == nil {
		line, err = context.stdin.ReadLine()
		if err == nil {
			_, err = fmt.Fprintln(context.stdout, line)
		}
	}
	return 0, err
}
