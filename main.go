package main

import (
	"fmt"

	"github.com/patelpreet422/wifi-file-transfer/cmd"
)

func main() {
	args, err := cmd.ParseCommandLineArgs()

	fmt.Printf("command: %v, port: %d, files: %v\n", args.Command, args.Port, args.Files)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("See you later, bye 👋.\n")
		return
	}
}
