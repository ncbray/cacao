package regenerate

import (
	"fmt"
	"github.com/ncbray/compilerutil/writer"
	"os"
)

func dirExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func writeHeader(packageName string, sourceFile string, out *writer.TabbedWriter) {
	out.WriteLine("package " + packageName)
	out.EndOfLine()
	out.WriteLine(fmt.Sprintf("/* Generated from %s, do not edit by hand. */", sourceFile))
	out.EndOfLine()
}
