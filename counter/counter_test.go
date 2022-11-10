package counter_test

import (
	"errors"
	"testing"

	"github.com/matt-tyler/go-tools-book/counter"
)

func TestCounterReturnsNextValue(t *testing.T) {
	t.Parallel()

	c, err := counter.New()
	if err != nil {
		t.Error(err)
	}

	for i := range [5]int{} {
		want := i
		got := c.Next()
		if got != want {
			t.Errorf("Want %d, got %d", want, got)
		}
	}
}

type ArrayBuffer struct {
	data []byte
	i    int
	cap  int
}

func New(cap int) *ArrayBuffer {
	i := 0
	data := make([]byte, cap)
	return &ArrayBuffer{data, i, cap}
}

func (b *ArrayBuffer) Write(p []byte) (n int, err error) {
	if b.i < b.cap {
		slice := b.data[b.i:b.cap]
		n = copy(slice, p)
		b.i += n
	}
	if n < len(p) {
		err = errors.New("could not write all data")
	}
	return
}

func (b *ArrayBuffer) String() string {
	return string(b.data)
}

func TestRunReturnsNextValueToWriter(t *testing.T) {
	t.Parallel()

	ch := make(chan struct{}, 7)
	for i := 0; i < 7; i += 1 {
		ch <- struct{}{}
	}

	defer close(ch)

	fakeTerminal := New(5)

	c, err := counter.New(counter.WithSignal(ch), counter.WithOutput(fakeTerminal))
	if err != nil {
		t.Error(err)
	}

	c.Run()

	want := "01234"
	got := fakeTerminal.String()
	if got != want {
		t.Errorf("Want %q, got %q", want, got)
	}
}
