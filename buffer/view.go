package buffer

import (
	"deedles.dev/axion/internal/util"
	"golang.org/x/exp/slices"
)

type View struct {
	buf           *Buffer
	start, length int
}

func (v *View) Cursor(line, col int) *Cursor {
	return &Cursor{
		view: v,
		loc:  util.Min(v.buf.lineIndex(line)+col, len(v.buf.data)),
	}
}

func (v *View) String() string {
	i1, i2 := v.buf.sliceByLines(v.start, v.length)
	return string(v.buf.data[i1:i2])
}

func (v *View) Runes() []rune {
	i1, i2 := v.buf.sliceByLines(v.start, v.length)
	return slices.Clone(v.buf.data[i1:i2])
}

func (v *View) Resize(length int) {
	v.length = length
}
