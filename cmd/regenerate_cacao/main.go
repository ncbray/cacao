package main

import (
	"encoding/json"
	"fmt"
	"github.com/ncbray/cacao/regenerate"
	"github.com/ncbray/compilerutil/fs"
	"io/ioutil"
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

func declarationsFromJSON(data_dir string) *regenerate.Declarations {
	decls := &regenerate.Declarations{}

	enums := jsonFilesInDir(filepath.Join(data_dir, "dsl/enum"))
	trees := jsonFilesInDir(filepath.Join(data_dir, "dsl/tree"))

	for _, filename := range enums {
		fmt.Println("Parsing", filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		decl := &regenerate.EnumDecl{}
		err = json.Unmarshal(data, decl)
		if err != nil {
			panic(err)
		}
		_, decl.DataSource = filepath.Split(filename)
		decls.Enums = append(decls.Enums, decl)
	}
	for _, filename := range trees {
		fmt.Println("Processing", filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		decl := &regenerate.TreeDecl{}
		err = json.Unmarshal(data, decl)
		if err != nil {
			panic(err)
		}
		_, decl.DataSource = filepath.Split(filename)
		decls.Trees = append(decls.Trees, decl)
	}
	return decls
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
	Synthesize(data_dir, declarations, fsys)

	output_dir := filepath.Join(gopath, "src")
	regenerate.ProcessDeclarations(declarations, output_dir, fsys)
	fsys.Commit()
}
