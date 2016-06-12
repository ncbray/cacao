package regenerate

import (
	"fmt"
	"github.com/ncbray/compilerutil/fs"
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

type Declarations struct {
	Enums []*EnumTypeDecl
	Trees []*TreeDecl
}

func ProcessDeclarations(decl *Declarations, output_dir string, fsys fs.FileSystem) {
	for _, e := range decl.Enums {
		ProcessEnum(e, output_dir, fsys)
	}
	for _, t := range decl.Trees {
		ProcessTree(t, output_dir, fsys)
	}
}
