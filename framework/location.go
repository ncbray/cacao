package framework

import (
	"unicode/utf8"
)

func findLines(data []byte) []int {
	lines := []int{0}
	for i, r := range data {
		if r == '\n' {
			lines = append(lines, i+1)
		}
	}
	lines = append(lines, len(data))
	return lines
}

type FileLocationProvider struct {
	filename string
	data     []byte
	lines    []int
}

func (loc *FileLocationProvider) GetLocationInfo(pos int) (string, int, int, string) {
	var line int
	for line = 1; line < len(loc.lines); line++ {
		if loc.lines[line] > pos {
			break
		}
	}
	begin := loc.lines[line-1]
	end := loc.lines[line]

	index := begin
	col := 0
	runes := []rune{}
	for index < end {
		if index <= pos {
			col = len(runes)
		}
		r, size := utf8.DecodeRune(loc.data[index:])
		if r == '\t' {
			for i := 0; i < 4; i++ {
				runes = append(runes, ' ')
			}
		} else {
			runes = append(runes, r)
		}
		index += size
	}

	// Strip off trailing newline.
	for len(runes) > 0 && (runes[len(runes)-1] == '\n' || runes[len(runes)-1] == '\r') {
		runes = runes[:len(runes)-1]
	}

	return loc.filename, line, col, string(runes)
}

func MakeFileLocationProvider(filename string, data []byte) *FileLocationProvider {
	return &FileLocationProvider{
		filename: filename,
		data:     data,
		lines:    findLines(data),
	}
}
