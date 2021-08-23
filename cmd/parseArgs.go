package cmd

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/jhoonb/archivex"
)

type ParsedArgs struct {
	Command Command
	Port    int
	Files   []string
}

func ZipFiles(files []string) (string, error) {
	zip := new(archivex.ZipFile)
	tmpfile, err := ioutil.TempFile("", "output")
	if err != nil {
		return "", err
	}
	tmpfile.Close()
	if err := os.Rename(tmpfile.Name(), tmpfile.Name()+".zip"); err != nil {
		return "", err
	}
	zip.Create(tmpfile.Name() + ".zip")
	for _, filename := range files {
		fileinfo, err := os.Stat(filename)
		if err != nil {
			return "", err
		}
		if fileinfo.IsDir() {
			zip.AddAll(filename, true)
		} else {
			file, err := os.Open(filename)
			if err != nil {
				return "", err
			}
			defer file.Close()
			if err := zip.Add(filename, file, fileinfo); err != nil {
				return "", err
			}
		}
	}
	if err := zip.Close(); err != nil {
		return "", nil
	}
	return zip.Name, nil
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
