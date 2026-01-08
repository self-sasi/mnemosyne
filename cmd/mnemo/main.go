package main

import (
	"github.com/self-sasi/mnemosyne/internal/cli"
	"github.com/self-sasi/mnemosyne/internal/logging"
)

func init() {
	logging.SetupBaseLogger()
}

func main() {
	cli.Execute()
}
