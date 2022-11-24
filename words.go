package count

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) Print() {
	fmt.Fprintln(p.Output, "Helllllo!")
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}

func Print() {
	NewPrinter().Print()
}

}
