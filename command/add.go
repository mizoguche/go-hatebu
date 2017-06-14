package command

import (
	"fmt"
	"os"

	"github.com/mizoguche/go-hatebu/hatebu"
)

type AddBookmarkCommand struct {
}

func (c *AddBookmarkCommand) Synopsis() string {
	return "Add new bookmark"
}

func (c *AddBookmarkCommand) Help() string {
	return "Usage: go-hatebu add <url> [<comment>]"
}

func (c *AddBookmarkCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Print("url is required\n")
		fmt.Print(c.Help() + "\n")
		return -1
	}

	url := args[0]
	var comment string
	if len(args) > 1 {
		comment = args[1]
	}

	postTwitter := os.Getenv("HATENA_POST_TWITTER")
	postFacebook := os.Getenv("HATENA_POST_FACEBOOK")
	postEvernote := os.Getenv("HATENA_POST_EVERNOTE")

	return hatebu.Add(hatebu.AddRequest{
		URL:          url,
		Comment:      comment,
		PostTwitter:  postTwitter != "",
		PostFacebook: postFacebook != "",
		PostEvernote: postEvernote != "",
	})
}
