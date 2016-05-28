package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

var (
	pkgLinePrefix = []byte("package ")
)

type errFileExists string

func (e errFileExists) Error() string {
	return "file exists: " + string(e)
}

func copy(dest, src, path, id, license string, files ...string) ([]string, error) {
	pkgName, err := getPkgName(dest)
	if err != nil {
		return nil, err
	}
	var copied []string
	for _, file := range files {
		destFile, err := copyfile(dest, src, path, file, id, pkgName, license)
		if err != nil {
			return nil, err
		}
		copied = append(copied, destFile)
	}
	return copied, nil
}

func copyfile(dest, src, path, file, id, pkgName, license string) (string, error) {
	fname := filepath.Base(file)
	destFile := filepath.Join(dest, fmt.Sprintf(filenameFormat, fname))
	destFilename := filepath.Base(destFile)
	info("cp", fname, destFilename)

	in, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer in.Close()

	if !forceOverwrite {
		// make sure we don't overwrite unless -f is set
		_, err = os.Stat(destFile)
		if !os.IsNotExist(err) /* file exists */ {
			return "", errFileExists(destFile)
		}
	}

	out, err := os.Create(destFile)
	if err != nil {
		return "", err
	}
	defer out.Close()

	err = writeHeader(out, src, path, id, license)
	if err != nil {
		return "", err
	}

	s := bufio.NewScanner(in)
	for s.Scan() {
		if bytes.HasPrefix(s.Bytes(), pkgLinePrefix) {
			fmt.Fprintln(out, "package", pkgName)
			continue
		}
		fmt.Fprintln(out, s.Text())
	}

	err = out.Sync()
	if err != nil {
		return "", err
	}

	on("+", destFilename)
	return destFile, nil
}
