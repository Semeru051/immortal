package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/starrysilk/immortal"
	"github.com/starrysilk/immortal/cmd/commands"
)

func main() {
	if len(os.Args) < 2 {
		commands.HandleHelp(os.Args)
		commands.ExitOnError(errors.New("at least 1 arguments expected"))
	}

	switch os.Args[1] {
	case "run":
		commands.HandleRun(os.Args)

	case "help":
		commands.HandleHelp(os.Args)
		os.Exit(0)

	case "version":
		fmt.Println(immortal.StringVersion()) //nolint
		os.Exit(0)

	default:
		commands.HandleHelp(os.Args)
	}
}
