package main

import (
	"os"

	"github.com/tyhal/protolint/internal/cmd"
)

func main() {
	os.Exit(int(
		cmd.Do(
			os.Args[1:],
			os.Stdout,
			os.Stderr,
		),
	))
}
