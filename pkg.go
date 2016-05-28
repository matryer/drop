package main

import (
	"errors"
	"go/parser"
	"go/token"
	"os"
)

// pkg gets the go package name from the specified directory.
func pkg(p string) (string, error) {
	fset := token.NewFileSet()
	first := true
	pkgs, err := parser.ParseDir(fset, p, func(os.FileInfo) bool {
		// process only a single file
		defer func() { first = false }()
		return first
	}, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	for k := range pkgs {
		return k, nil // return first package
	}
	return "", errors.New("failed to discover package")
}
