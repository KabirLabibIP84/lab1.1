package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"

	"github.com/KabirLabibIP84/lab1.1/tree/master/build/gomodule"
	"github.com/KabirLabibIP84/lab1.1/blob/master/build/gomodule/zip_archive.go"
)

var (
	dryRun  = flag.Bool("dry-run", false, "Generate ninja build file but don't start the build")
	verbose = flag.Bool("v", false, "Display debugging logs")
)

func NewContext() *blueprint.Context {
	ctx := bood.PrepareContext()
	ctx.RegisterModuleType("go_testedbinary", gomodule.SimpleBinFactory)
	ctx.RegisterModuleType("zip_archive", zip_archive.SimpleZipArchiveFactory)
	return ctx
}

func main() {
	flag.Parse()

	config := bood.NewConfig()
	if !*verbose {
		config.Debug = log.New(ioutil.Discard, "", 0)
	}
	ctx := NewContext()

	ninjaBuildPath := bood.GenerateBuildFile(config, ctx)

	if !*dryRun {
		config.Info.Println("Starting the build now")

		cmd := exec.Command("ninja", append([]string{"-f", ninjaBuildPath}, flag.Args()...)...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			config.Info.Fatal("Error invoking ninja build. See logs above.")
		}
	}
}
