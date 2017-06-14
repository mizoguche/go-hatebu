package command

import "github.com/mizoguche/go-hatebu/hatebu"

type AddBookmarkCommand struct {
}

func (c *AddBookmarkCommand) Synopsis() string {
	return "Add new bookmark"
}

func (c *AddBookmarkCommand) Help() string {
	return "Usage: go-hatebu add url comment"
}

func (c *AddBookmarkCommand) Run(args []string) int {
	hatebu.Add("", "")
	return 0
}
