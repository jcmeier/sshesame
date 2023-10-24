package commands

type cmdTrue struct{}

func (cmdTrue) execute(context CommandContext) (uint32, error) {
	return 0, nil
}

type cmdFalse struct{}

func (cmdFalse) execute(context CommandContext) (uint32, error) {
	return 1, nil
}
