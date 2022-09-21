package buffer

import (
	"bytes"
	"fmt"

	"golang.org/x/exp/slices"
)

type Cursor struct {
	buf *Buffer
	loc int
}

func (c *Cursor) Location() int {
	return c.loc
}

func (c *Cursor) LineAndCol() (line, col int) {
	for i, start := range c.buf.lines {
		if c.loc < start {
			fmt.Println(c.loc - c.buf.lines[i-1])
			return i - 1, c.loc - c.buf.lines[i-1]
		}
	}
	return 0, 0
}

// Write writes data into the buffer.
func (c *Cursor) Write(data []byte) (int, error) {
	fmt.Println(c.loc)
	r := bytes.Runes(data)
	c.buf.data = slices.Insert(c.buf.data, c.loc, r...)
	c.buf.updateLines()
	c.loc += len(r)
	return len(data), nil
}

func (c *Cursor) WriteByte(char byte) error {
	c.buf.data = slices.Insert(c.buf.data, c.loc, rune(char))
	c.buf.updateLines()
	c.loc++
	return nil
}
