package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/mizoguche/go-hatebu/hatebu"
)

func main() {
	os.Exit(Run(os.Args))
}

// Run command
func Run(args []string) int {
	c := cli.NewCLI("go-gatevy", "0.1.0")
	c.Args = args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &hatebu.AddBookmarkCommand{}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to execute: %s\n", err.Error())
	}
	return exitCode
}
