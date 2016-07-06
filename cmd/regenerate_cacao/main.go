package main

import (
	"github.com/ncbray/cacao/irgen"
	"github.com/ncbray/compilerutil/fs"
	"os"
	"path/filepath"
	"strings"
)

func jsonFilesInDir(gopath string) []string {
	dirname := filepath.Join(gopath, "src/github.com/ncbray/cacao/data")

	f, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}
	names, err := f.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	out := []string{}
	for _, name := range names {
		if strings.HasSuffix(name, ".json") {
			out = append(out, filepath.Join(dirname, name))
		}
	}
	return out
}

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		println("GOPATH must be set.")
	}

	tdir, err := fs.MakeTempDir("compiler_outputs_")
	if err != nil {
		panic(err)
	}
	defer tdir.Cleanup()

	fsys := fs.MakeBufferedFileSystem(tdir)

	data_dir := filepath.Join(gopath, "src/github.com/ncbray/cacao/data")
	decls := Synthesize(data_dir, []*irgen.Declarations{ast, struct_ir, graph}, fsys)

	output_dir := filepath.Join(gopath, "src")
	irgen.ProcessDeclarations(decls, output_dir, fsys)
	fsys.Commit()
}
