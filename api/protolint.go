package api

import (
	"errors"
	"github.com/tyhal/protolint/internal/cmd"
	"io"
	"strconv"
)

// Lint allows tyhal to use this thing as a lib
func Lint(file string,
	fix bool,
	stdout io.Writer,
	stderr io.Writer) (err error) {
	args := []string{"lint"}
	if fix {
		args = append(args, "-fix")
	}
	exit := cmd.Do(
		append(args, file),
		stdout,
		stderr,
	)
	if exit != 0 {
		err = errors.New("protolint returned: " + strconv.Itoa(int(exit)))
	}
	return
}
