package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	getPkgName     = pkg
	filenameFormat = "%s"
	forceOverwrite = false
)

func noop() {}

func main() {
	var (
		pkg         = flag.String("package", "", "package name (default auto discover)")
		verbose     = flag.Bool("v", false, "verbose logging")
		help        = flag.Bool("help", false, "show help")
		showVersion = flag.Bool("version", false, "print version")
		outformat   = flag.String("outformat", "%s", "filename format")
		force       = flag.Bool("f", false, "overwrite files")
	)
	flag.Parse()
	if *help {
		usage()
		return
	}
	if *showVersion {
		fmt.Println(version)
		return
	}
	if len(*pkg) > 0 {
		getPkgName = func(string) (string, error) {
			return *pkg, nil
		}
	}
	if *verbose {
		info = on
	}
	if len(*outformat) > 0 {
		filenameFormat = *outformat
	}
	forceOverwrite = *force
	args := flag.Args()
	if len(args) < 1 {
		usage()
		fatal("must provide source")
	}
	if len(args) > 2 {
		usage()
		fatal("wrong number of arguments")
	}
	src := args[0]
	path := ""
	dest := "."
	if len(args) > 1 {
		path = args[1]
	}

	files, err := gimme(src, path, dest)
	if err != nil {
		fatal(err)
	}

	info("copied", len(files), "file(s)")
}

func fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}

func usage() {
	fmt.Println(`gimme v` + version + `
https://github.com/matryer/gimme

usage:
  gimme [flags] import [path]

  flags       - see below
  import      - import path to go get
  path        - directory within repo to copy from

flags:`)
	flag.PrintDefaults()
}
