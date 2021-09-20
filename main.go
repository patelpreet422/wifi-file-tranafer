package main

import (
	"fmt"

	"github.com/patelpreet422/wifi-file-transfer/cmd"
	"github.com/patelpreet422/wifi-file-transfer/util"
)

func main() {
	args, err := cmd.ParseCommandLineArgs()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if args.Command == cmd.NewCommandRegistry().Send {
		payload, err := util.GetPayloadFromArgs(args.Files)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		fmt.Printf("payload: %v\n", payload)
	}

}
