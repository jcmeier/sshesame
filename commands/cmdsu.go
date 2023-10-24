package commands

type cmdSu struct{}

func (cmdSu) execute(context CommandContext) (uint32, error) {
	newContext := context
	newContext.user = "root"
	if len(context.args) > 1 {
		newContext.user = context.args[1]
	}
	newContext.args = ShellProgram
	return ExecuteProgram(newContext)
}
