package exec

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/ilackarms/pkg/errors"
)

type command struct {
	*exec.Cmd
}

func (cmd *command) Dir(dir string) *command {
	cmd.Cmd.Dir = dir
	return cmd
}

func (cmd *command) Run() error {
	fmt.Fprintf(cmd.Stdout, "running command: %v\n", cmd.Args)
	if err := cmd.Cmd.Run(); err != nil {
		return errors.New("failed running command "+strings.Join(cmd.Args, " "), err)
	}
	return nil
}

func Command(out io.Writer, args ...string) *command {
	cmd := exec.Command(args[0])
	cmd.Args = args
	cmd.Stdout = out
	cmd.Stderr = out
	return &command{cmd}
}

func RunCommand(out io.Writer, args ...string) error {
	return Command(out, args...).Run()
}
