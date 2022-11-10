package counter

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int

	out io.Writer

	ch <-chan struct{}
}

type CounterOption = func(c *Counter) error

func WithSignal(ch <-chan struct{}) CounterOption {
	return func(c *Counter) error {
		c.ch = ch
		return nil
	}
}

func WithOutput(w io.Writer) CounterOption {
	return func(c *Counter) error {
		c.out = w
		return nil
	}
}

func New(options ...CounterOption) (c *Counter, err error) {
	c = &Counter{value: -1}
	for _, option := range options {
		err = option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, err
}

func (c *Counter) Set(v int) {
	c.Lock()
	defer c.Unlock()

	c.value = v
}

func (c *Counter) Next() int {
	c.Lock()
	defer c.Unlock()

	c.value += 1
	return c.value
}

func (c *Counter) Run() (err error) {

	ticker := time.NewTicker(1 * time.Second)
	if c.ch != nil {
		ticker.Stop()
	}

	for err == nil {
		select {
		case <-ticker.C:
		case _, ok := <-c.ch:
			if !ok {
				return errors.New("channel is closed")
			}
		}
		next := c.Next()
		_, err = fmt.Fprintf(c.out, "%d", next)
	}
	return err
}
