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
	showCommand bool
}

func (cmd *command) ShowCommand(enable bool) *command {
	cmd.showCommand = enable
	return cmd
}

func (cmd *command) Dir(dir string) *command {
	cmd.Cmd.Dir = dir
	return cmd
}

func (cmd *command) Run() error {
	if cmd.showCommand {
		fmt.Fprintf(cmd.Stdout, "running command: %v\n", cmd.Args)
	}
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
	return &command{
		Cmd:         cmd,
		showCommand: true,
	}
}

func RunCommand(out io.Writer, args ...string) error {
	return Command(out, args...).Run()
}
