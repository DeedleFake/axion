// Package buffer implements a data buffer useful for editing UTF-8 encoded text.
package buffer

// A Buffer holds data for manipulation. Its zero value is ready to
// use.
type Buffer struct {
	data []rune
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
	i1 = b.lineIndex(start)

	line := start + length
	for i, c := range b.data[i1:] {
		if line == 0 {
			return i1, i
		}

		if c == '\n' {
			line--
		}
	}

	return i1, len(b.data)
}

func (b *Buffer) lineIndex(line int) int {
	for i, c := range b.data {
		if line == 0 {
			return i
		}

		if c == '\n' {
			line--
		}
	}

	return len(b.data)
}

func (b *Buffer) lineOfIndex(i int) (line, beginning int) {
	if i >= len(b.data) {
		i = len(b.data) - 1
	}

	for ; i >= 0; i-- {
		if b.data[i] == '\n' {
			line++
			if beginning == 0 {
				beginning = i + 1
			}
		}
	}
	return line, beginning
}
