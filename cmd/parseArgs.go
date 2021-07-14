package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/fatih/color"
)

type ParsedArgs struct {
	Command Command
	Port    int
	Files   []string
}

func ParseCommandLineArgs() (ParsedArgs, error) {
	defaultPort := 65432
	port := flag.Int("p", defaultPort, fmt.Sprintf("port on which server will run, if unspecified then %d port will be used", defaultPort))
	flag.Parse()

	Commands := NewCommandRegistry()

	args := flag.Args()

	if len(args) == 0 {
		errorMessage := fmt.Sprintf("no argument passed, pass either %v to send the file or %v to receive file", color.GreenString("send"), color.GreenString("receive"))
		return ParsedArgs{Command: Commands.None, Port: -1}, errors.New(errorMessage)
	}

	if args[0] != "send" && args[0] != "receive" {
		errorMessage := fmt.Sprintf("command not identified, pass either %v to send the file or %v to receive file", color.GreenString("send"), color.GreenString("receive"))
		return ParsedArgs{Command: Commands.None, Port: -1}, errors.New(errorMessage)
	}

	if args[0] == "send" && len(args) == 1 {
		return ParsedArgs{Command: Commands.None, Port: -1}, errors.New("choose atleast one file to send")
	}

	if args[0] == "send" {
		return ParsedArgs{Command: Commands.Send, Port: *port, Files: args[1:]}, nil
	}

	return ParsedArgs{Command: Commands.Receive, Port: *port}, nil
}
