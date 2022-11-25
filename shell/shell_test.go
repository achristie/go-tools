package shell_test

import (
	"count/shell"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString(t *testing.T) {
	t.Parallel()

	input := "/bin/ls -l main.go"
	want := []string{"/bin/ls", "-l", "main.go"}

	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromEmptyString(t *testing.T) {
	t.Parallel()

	input := ""
	_, err := shell.CmdFromString(input)
	if err == nil {
		t.Fatal("want error, got nil")
	}
}
