package pipeline

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Output io.Writer
	Error  error
	Reader io.Reader
}

func (p *Pipeline) Column(i int) *Pipeline {
	if p.Error != nil {
		p.Reader = strings.NewReader("")
	}
	if i <= 0 {
		p.Error = errors.New("index must be greater than 0")
		return p
	}

	result := &bytes.Buffer{}
	scanner := bufio.NewScanner(p.Reader)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < i {
			continue
		}
		fmt.Fprintln(result, fields[i-1])
	}

	return &Pipeline{Reader: result}
}

func (p *Pipeline) Freq() *Pipeline {
	return p
}

func FromString(s string) *Pipeline {
	p := New()
	p.Reader = strings.NewReader(s)
	return p
}

func FromFile(path string) *Pipeline {
	f, err := os.Open(path)
	if err != nil {
		return &Pipeline{Error: err}
	}
	p := New()
	p.Reader = f
	return p
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}

	io.Copy(p.Output, p.Reader)
}

func New() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

func (p *Pipeline) String() (string, error) {
	if p.Error != nil {
		return "", p.Error
	}
	b, err := io.ReadAll(p.Reader)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
