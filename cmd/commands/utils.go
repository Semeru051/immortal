package commands

import (
	"os"

	"github.com/starrysilk/immortal/pkg/logger"
)

func ExitOnError(err error) {
	logger.Error("immortal error", "err", err.Error())
	os.Exit(1)
}
