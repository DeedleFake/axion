// Package buffer implements a data buffer useful for editing UTF-8 encoded text.
package buffer

import (
	"bytes"

	"deedles.dev/axion/internal/util"
)

// A Buffer holds data for manipulation.
type Buffer struct {
	data  []rune
	lines []int
}

// New returns a new empty buffer.
func New() *Buffer {
	return &Buffer{}
}

func (b *Buffer) updateLines() {
	b.lines = b.lines[:0]
	b.lines = append(b.lines, 0)
	for i, c := range b.data {
		if c != '\n' {
			continue
		}

		b.lines = append(b.lines, i+1)
	}
}

// Write writes data into the buffer.
func (b *Buffer) Write(data []byte) (int, error) {
	b.data = append(b.data, bytes.Runes(data)...)
	b.updateLines()
	return len(data), nil
}

// View returns a slice of the data in the buffer starting at the
// beginning of the line numbered start and containing up to length
// lines. The caller should not modify the returned slice nor should
// they hold onto it after calling any methods that manipulate the
// data in the buffer.
//
// Note that line numbers begin at 0, not 1.
func (b *Buffer) View(start, length int) []rune {
	si, ei := b.sliceByLines(start, length)
	return b.data[si:ei:ei]
}

// sliceByLines returns the start and end indices into b.data that
// begins at the line with the given index and is at least length
// lines long.
func (b *Buffer) sliceByLines(start, length int) (i1, i2 int) {
	if start+length >= len(b.lines) {
		start = len(b.lines) - length
	}
	if start < 0 {
		start = 0
	}

	return b.lines[start], b.lines[util.Min(start+length, len(b.lines)-1)]
}
