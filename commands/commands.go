package commands

import (
	"fmt"
	"io"
)

type ReadLiner interface {
	ReadLine() (string, error)
}

type CommandContext struct {
	args           []string
	stdin          ReadLiner
	stdout, stderr io.Writer
	pty            bool
	user           string
}

func CreateCommandContext(
	args []string,
	stdin ReadLiner,
	stdout, stderr io.Writer,
	pty bool,
	user string,

) CommandContext {
	return CommandContext{
		args, stdin, stdout, stderr, pty, user}
}

type command interface {
	execute(context CommandContext) (uint32, error)
}

var commands = map[string]command{
	"sh":    cmdShell{},
	"true":  cmdTrue{},
	"false": cmdFalse{},
	"echo":  cmdEcho{},
	"cat":   cmdCat{},
	"su":    cmdSu{},
	"ls":    cmdLs{},
}

var ShellProgram = []string{"sh"}

func ExecuteProgram(context CommandContext) (uint32, error) {
	if len(context.args) == 0 {
		return 0, nil
	}
	command := commands[context.args[0]]
	if command == nil {
		_, err := fmt.Fprintf(context.stderr, "%v: command not found\n", context.args[0])
		return 127, err
	}
	return command.execute(context)
}
