package ui

import (
	"github.com/muesli/ansi"
)

func AnsiTrim(content string, width int) string {
	var out []rune
	seq := false
	size := 0
	for _, c := range content {
		if c == ansi.Marker {
			seq = true
			out = append(out, c)
		} else if seq {
			if ansi.IsTerminator(c) {
				seq = false
			}
			out = append(out, c)
		} else {
			if size < width {
				out = append(out, c)
				size++
			}
		}
	}

	return string(out)
}
