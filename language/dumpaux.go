package language

import (
	"github.com/ncbray/compilerutil/writer"
	"os"
)

func dumpPrefixOperator(value PrefixOperator, out *writer.TabbedWriter) {
	out.WriteString(value.String())
}

func dumpInfixOperator(value InfixOperator, out *writer.TabbedWriter) {
	out.WriteString(value.String())
}

func DumpFile(file *File) {
	out := writer.MakeTabbedWriter("  ", os.Stdout)
	dumpFile(file, out)
	out.EndOfLine()
}
