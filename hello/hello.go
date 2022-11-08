package hello

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}

func Print() {
	NewPrinter().Print()
}

func (p *Printer) Print() {
	fmt.Fprintln(p.Output, "Hello, World")
}
