package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		exp   int
	}{
		{name: "empty", input: "", exp: 0},
		{name: "single line", input: "AnDrEWW!", exp: 1},
		{name: "multi line", input: "This has a line\nbreak", exp: 2},
		{name: "repeated line breaks", input: "This \nhas a line\n\nbreak\n", exp: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := hello.NewCounter(
				hello.WithInput(bytes.NewBufferString(tt.input)),
			)

			if err != nil {
				t.Fatal(err)
			}
			got := c.Lines()
			if got != tt.exp {
				t.Errorf("want %d, got %d", tt.exp, got)
			}
		})

	}
}
