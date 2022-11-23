package search

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type searcher struct {
	input           io.Reader
	output          io.Writer
	caseInsensitive bool
}

type SearchOption = func(s *searcher) error

func WithCaseInsensitiveSearch() SearchOption {
	return func(s *searcher) error {
		s.caseInsensitive = true
		return nil
	}
}

func WithInput(r io.Reader) SearchOption {
	return func(s *searcher) error {
		s.input = r
		return nil
	}
}

func WithOutput(w io.Writer) SearchOption {
	return func(s *searcher) error {
		s.output = w
		return nil
	}
}

func New(options ...SearchOption) (*searcher, error) {
	s := &searcher{input: os.Stdin, output: os.Stdout}

	for _, option := range options {
		err := option(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (s searcher) Lines(query string) error {
	scanner := bufio.NewScanner(s.input)

	substringIn := func(substr string) func(a string) bool {
		sub := substr
		if s.caseInsensitive {
			sub = strings.ToLower(sub)
		}
		return func(a string) bool {
			return strings.Contains(a, sub)
		}
	}(query)

	for scanner.Scan() {
		line := scanner.Text()
		if substringIn(line) {
			output := fmt.Sprintln(line)
			s.output.Write([]byte(output))
		}
	}
	return nil
}

func Lines(query string) error {
	s, err := New(WithInput(os.Stdin), WithOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	return s.Lines(query)
}
