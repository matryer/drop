package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultPattern = "*.go"
	gimmeFile      = ".gimme"
)

var ignoreFiles = []string{"doc.go"}

var errNeedDir = errors.New("expected directory")

// files gets a list of files to include from the specified path.
// If no .gimme file is specified, default files are selected.
func files(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errNeedDir
	}

	files, err := gimmefile(path)
	if err == nil {
		return files, nil
	}
	files, err = filepath.Glob(filepath.Join(path, defaultPattern))
	if err != nil {
		return nil, err
	}
	return cleanFiles(files...), nil
}

type errLine struct {
	file string
	n    int
	err  error
}

func (e errLine) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.file, e.n, e.err)
}

// license gets the path to a license file (if there is one).
func license(p string) (string, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		base := filepath.Base(file.Name())
		base = strings.TrimSuffix(base, filepath.Ext(base))
		if strings.ToLower(base) == "license" {
			return extractFirstLine(filepath.Join(p, file.Name()))
		}
	}
	return "", nil
}

func extractFirstLine(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		clean := strings.TrimSpace(s.Text())
		if len(clean) > 0 {
			return clean, nil
		}
	}
	return "", nil
}

// gimmefile gets a cleaned list of files from the .gimme
// file at the specified path.
func gimmefile(path string) ([]string, error) {
	gimmefile := filepath.Join(path, gimmeFile)
	src, err := ioutil.ReadFile(gimmefile)
	if err != nil {
		return nil, err
	}
	var files []string
	s := bufio.NewScanner(bytes.NewReader(src))
	lineNumber := 0
	for s.Scan() {
		lineNumber++
		text := s.Text()
		if strings.HasPrefix(text, "#") || len(text) == 0 {
			// ignore comments and empty lines
			continue
		}
		relPattern := filepath.Join(path, text)
		theseFiles, err := filepath.Glob(relPattern)
		if err != nil {
			return nil, errLine{file: gimmefile, n: lineNumber, err: err}
		}
		files = append(files, theseFiles...)
	}
	return cleanFiles(files...), nil
}

// cleanFiles ignores ignoreFiles and selects only unique
// files from the list.
func cleanFiles(files ...string) []string {
	var cleanFiles []string
	for _, file := range files {
		include := true
		for _, existing := range cleanFiles {
			if existing == file {
				include = false
				break
			}
		}
		for _, ignore := range ignoreFiles {
			match, err := filepath.Match(ignore, filepath.Base(file))
			if err != nil {
				log.Println("gimme: filepath.Match:", err)
			}
			if match {
				include = false
				break
			}
		}
		if include {
			cleanFiles = append(cleanFiles, file)
		}
	}
	return cleanFiles
}
