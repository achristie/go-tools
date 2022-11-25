package pipeline_test

import (
	"bytes"
	"count/pipeline"
	"errors"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStdOut(t *testing.T) {
	t.Parallel()

	want := "hello\n"
	p := pipeline.FromString(want)
	buf := &bytes.Buffer{}

	p.Output = buf
	p.Stdout()

	if p.Error != nil {
		t.Fatal(p.Error)
	}

	got := buf.String()

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStdOutError(t *testing.T) {
	t.Parallel()

	want := "hello\n"
	p := pipeline.FromString(want)
	p.Error = errors.New("err!")

	buf := &bytes.Buffer{}
	p.Output = buf
	p.Stdout()
	got := buf.String()

	if got != "" {
		t.Errorf("want nothing on error, got %q", got)
	}

}

func TestFromFile(t *testing.T) {
	t.Parallel()

	want := []byte("hello\n")
	p := pipeline.FromFile("testdata/hello.txt")
	if p.Error != nil {
		t.Fatal(p.Error)
	}

	got, err := io.ReadAll(p.Reader)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(got, want) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFromInvalidFile(t *testing.T) {
	t.Parallel()

	p := pipeline.FromFile("testdata/invalid.txt")
	if p.Error == nil {
		t.Error("want error, but did not get one")
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	want := "hello\n"
	p := pipeline.FromString(want)

	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("hello\n")
	p.Error = errors.New("err")
	_, err := p.String()

	if err == nil {
		t.Error("want error, got nil")
	}
}

func TestColumn(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("1 2 3\n1 2 3\n1 2 3\n")
	want := "2\n2\n2\n"

	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestColumnError(t *testing.T) {
	t.Parallel()

	want := "1 2 3\n"
	p := pipeline.FromString(want)
	p.Error = errors.New("err!")

	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) > 0 {
		t.Errorf("want no output from Column after error, got %q", data)
	}
}

func TestColumnInvalid(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("")
	p.Column(-1)
	if p.Error == nil {
		t.Error("want error on non-positive Column, but got nil")
	}
}
