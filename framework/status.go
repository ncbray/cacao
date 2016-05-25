package framework

import (
	"fmt"
	"io"
)

type CompileStatus interface {
	Error(text string)
	LocationError(loc int, text string)
	ErrorOccured() bool
}

type compileStatus struct {
	loc     *FileLocationProvider
	out     io.Writer
	errored bool
}

func (status *compileStatus) Error(message string) {
	io.WriteString(status.out, "ERROR ")
	io.WriteString(status.out, message)
	io.WriteString(status.out, "\n")
	status.errored = true
}

func (status *compileStatus) LocationError(loc int, message string) {
	filename, lineNum, colNum, lineText := status.loc.GetLocationInfo(loc)
	io.WriteString(status.out, fmt.Sprintf("ERROR %s:%d:%d: %s\n", filename, lineNum, colNum, message))
	io.WriteString(status.out, lineText)
	io.WriteString(status.out, "\n")
	for i := 0; i < colNum; i++ {
		io.WriteString(status.out, " ")
	}
	io.WriteString(status.out, "^\n")
	status.errored = true
}

func (status *compileStatus) ErrorOccured() bool {
	return status.errored
}

func MakeCompileStatus(loc *FileLocationProvider, out io.Writer) CompileStatus {
	return &compileStatus{loc: loc, out: out}
}
