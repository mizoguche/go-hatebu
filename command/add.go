package command

type AddBookmarkCommand struct {
}

func (c *AddBookmarkCommand) Synopsis() string {
	return "Add new bookmark"
}

func (c *AddBookmarkCommand) Help() string {
	return "Usage: go-hatebu add url comment"
}

func (c *AddBookmarkCommand) Run(args []string) int {
	// TODO: Implement
	return 0
}
