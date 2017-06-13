package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/mizoguche/go-hatebu/command"
)

func main() {
	os.Exit(Run(os.Args))
}

// Run command
func Run(args []string) int {
	c := cli.NewCLI("go-hatebu", "0.1.0")
	c.Args = args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &command.AddBookmarkCommand{}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to execute: %s\n", err.Error())
	}
	return exitCode
}
