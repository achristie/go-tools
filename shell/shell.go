package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	DryRun bool
}

func CmdFromString(s string) (*exec.Cmd, error) {
	args := strings.Fields(s)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}

	return exec.Command(args[0], args[1:]...), nil
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
}

func RunCLI() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	output := s.Stdout
	for {
		fmt.Fprint(output, "> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(output, "\nBye!")
			break
		}

		cmd, err := CmdFromString(line)
		if err != nil {
			continue
		}

		if s.DryRun {
			fmt.Fprintf(output, "%s", line)
			continue
		}

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(output, "error:", err)
		}

		fmt.Fprintf(output, "%s", out)
	}
}
