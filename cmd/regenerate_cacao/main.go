package main

import (
	"github.com/ncbray/cacao/regenerate"
	"github.com/ncbray/compilerutil/fs"
	"os"
	"path/filepath"
	"strings"
)

func jsonFilesInDir(dirname string) []string {
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

	data_dir := filepath.Join(gopath, "src/github.com/ncbray/cacao/data")
	enums := jsonFilesInDir(filepath.Join(data_dir, "dsl/enum"))
	trees := jsonFilesInDir(filepath.Join(data_dir, "dsl/tree"))

	tdir, err := fs.MakeTempDir("compiler_outputs_")
	if err != nil {
		panic(err)
	}
	defer tdir.Cleanup()

	fsys := fs.MakeBufferedFileSystem(tdir)
	language_dir := filepath.Join(gopath, "src")
	for _, filename := range enums {
		regenerate.ProcessEnumFile(filename, language_dir, fsys)
	}
	for _, filename := range trees {
		regenerate.ProcessTreeFile(filename, language_dir, fsys)
	}
	fsys.Commit()
}
