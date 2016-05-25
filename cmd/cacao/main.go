package main

import (
	"flag"
	"fmt"
	"github.com/ncbray/cacao/framework"
	"github.com/ncbray/cacao/language"
	"github.com/ncbray/cmdline"
	"io/ioutil"
	"os"
)

type CacaoFlags struct {
	InputFile string
}

func flagError(message string) {
	fmt.Println(message)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var config CacaoFlags

	app := cmdline.MakeApp("cacao")
	app.RequiredArgs([]*cmdline.Argument{
		{
			Name:  "input_file",
			Value: cmdline.String.Set(&config.InputFile),
		},
	})
	app.Run(os.Args[1:])

	filename := config.InputFile

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	loc := framework.MakeFileLocationProvider(filename, data)
	status := framework.MakeCompileStatus(loc, os.Stdout)

	file, deepest, ok := language.Parse(data)
	if !ok || deepest < len(data) {
		status.LocationError(deepest, "unexpected character")
	}
	if status.ErrorOccured() {
		os.Exit(1)
	}

	language.DumpFile(file)

	language.SemanticPass(file, status)

	if status.ErrorOccured() {
		os.Exit(1)
	}
}
