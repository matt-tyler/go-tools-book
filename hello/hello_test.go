package hello_test

import (
	"bytes"
	"testing"

	"github.com/matt-tyler/go-tools-book/hello"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}

	p := &hello.Printer{
		Output: fakeTerminal,
	}

	p.Print()

	want := "Hello, World\n"
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
