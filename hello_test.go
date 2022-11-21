package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func Test_Hello(t *testing.T) {
	t.Parallel()
	want := "Helllllo!\n"

	b := &bytes.Buffer{}
	p := hello.Printer{
		Output: b,
	}
	p.Print()
	got := b.String()

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
