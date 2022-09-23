package buffer

import (
	"bytes"

	"deedles.dev/axion/internal/util"
	"golang.org/x/exp/slices"
)

type Cursor struct {
	view *View
	loc  int
}

func (c *Cursor) Location() int {
	return c.loc
}

func (c *Cursor) Move(by int) {
	c.loc = util.Contain(0, len(c.view.buf.data), c.loc+by)
}

func (c *Cursor) MoveTo(i int) {
	c.loc = util.Contain(0, len(c.view.buf.data), i)
}

func (c *Cursor) MoveLines(num int) {
	if num < 0 {
		c.moveLinesUp(-num)
		return
	}
	c.moveLinesDown(num)
}

func (c *Cursor) moveLinesUp(num int) {
	data := c.view.buf.data

	lineend := c.loc
	col := -1

	for (lineend >= 0) && (num >= 0) {
		lineend = util.LastIndex(data[:lineend], '\n')
		if col < 0 {
			col = c.loc - lineend
		}
		num--
	}

	linelen := util.MaxIndex(data[lineend+1:], '\n') + 1
	c.loc = lineend + util.Min(col, linelen)
}

func (c *Cursor) moveLinesDown(num int) {
	data := c.view.buf.data

	linestart := util.LastIndex(data[:c.loc], '\n') + 1
	col := c.loc - linestart

	for (linestart < len(data)) && (num > 0) {
		linestart = linestart + util.MaxIndex(data[linestart:], '\n') + 1
		num--
	}
	if linestart >= len(data) {
		return
	}

	linelen := util.MaxIndex(data[linestart:], '\n')
	c.loc = linestart + util.Min(col, linelen)
}

func (c *Cursor) LineAndCol() (line, col int) {
	line, beginning := c.view.buf.lineOfIndex(c.loc)
	return line, c.loc - beginning
}

// Write writes data into the buffer.
func (c *Cursor) Write(data []byte) (int, error) {
	r := bytes.Runes(data)
	c.view.buf.data = slices.Insert(c.view.buf.data, c.loc, r...)
	c.loc += len(r)
	return len(data), nil
}

func (c *Cursor) WriteByte(char byte) error {
	c.view.buf.data = slices.Insert(c.view.buf.data, c.loc, rune(char))
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
	}
	if start < 0 {
		start = 0
	}
	if end > len(c.view.buf.data) {
		end = len(c.view.buf.data)
	}

	c.loc = start

	before := len(c.view.buf.data)
	c.view.buf.data = slices.Delete(c.view.buf.data, start, end)
	return before - len(c.view.buf.data)
}
