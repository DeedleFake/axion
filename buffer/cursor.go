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

func (c *Cursor) Move(by int) {
	c.loc = util.Contain(0, len(c.buf.data), c.loc+by)
}

func (c *Cursor) MoveTo(i int) {
	c.loc = util.Contain(0, len(c.buf.data), i)
}

func (c *Cursor) MoveLines(num int) {
	if num < 0 {
		c.moveLinesUp(-num)
		return
	}
	c.moveLinesDown(num)
}

func (c *Cursor) moveLinesUp(num int) {
	// TODO: There has got to be a better way to do this...
	var col int
	for i := c.loc - 1; i >= 0; i-- {
		if num == 0 {
			c.loc = i + col
			for i2 := i + 1; i2 < c.loc; i2++ {
				if c.buf.data[i2] == '\n' {
					c.loc = i2
					break
				}
			}
			return
		}

		if c.buf.data[i] == '\n' {
			num--
			if col == 0 {
				col = c.loc - (i + 1)
			}
		}
	}
}

func (c *Cursor) moveLinesDown(num int) {
	panic("Not implemented.")
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

// Delete deletes n characters starting at the cursor's current
// location. If n is negative, it deletes backwards. The final cursor
// location is the index of the first character deleted.
func (c *Cursor) Delete(n int) int {
	start := c.loc
	end := c.loc + n
	if end < start {
		start, end = end, start
		c.loc = start
	}

	before := len(c.buf.data)
	c.buf.data = slices.Delete(c.buf.data, start, end)
	return before - len(c.buf.data)
}
