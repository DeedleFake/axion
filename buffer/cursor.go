package buffer

import (
	"bytes"

	"deedles.dev/axion/internal/util"
	"golang.org/x/exp/slices"
)

type Cursor struct {
	buf *Buffer
	loc int
}

func (b *Buffer) Cursor(line, col int) *Cursor {
	return &Cursor{
		buf: b,
		loc: util.Min(b.lineIndex(line)+col, len(b.data)),
	}
}

func (c *Cursor) Location() int {
	return c.loc
}

func (c *Cursor) LineAndCol() (line, col int) {
	line, beginning := c.buf.lineOfIndex(c.loc)
	return line, c.loc - beginning
}

// Write writes data into the buffer.
func (c *Cursor) Write(data []byte) (int, error) {
	r := bytes.Runes(data)
	c.buf.data = slices.Insert(c.buf.data, c.loc, r...)
	c.loc += len(r)
	return len(data), nil
}

func (c *Cursor) WriteByte(char byte) error {
	c.buf.data = slices.Insert(c.buf.data, c.loc, rune(char))
	c.loc++
	return nil
}
