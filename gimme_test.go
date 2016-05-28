package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGimme(t *testing.T) {
	info = off

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "gimme-project")
	err := os.MkdirAll(dest, 0777)
	if err != nil {
		t.Errorf("MkDirAll: %s %v: %s", dest, 0777, err)
		return
	}
	defer os.RemoveAll(dest)
	// create a .go file with a different package name to 'try'
	err = ioutil.WriteFile(
		filepath.Join(dest, "code.go"), []byte(`package target
            func hi() {}`),
		0777)
	if err != nil {
		t.Errorf("WriteFile: %s", err)
		return
	}

	files, err := gimme("github.com/matryer/gimme-test", "", dest)
	if err != nil {
		t.Errorf("gimme: %s", err)
		return
	}
	if len(files) != 2 {
		t.Errorf("expected 2 files, but got %d", len(files))
	}

}

func TestGimmeSubPackages(t *testing.T) {
	info = off

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "gimme-project")
	err := os.MkdirAll(dest, 0777)
	if err != nil {
		t.Errorf("MkDirAll: %s %v: %s", dest, 0777, err)
		return
	}
	defer os.RemoveAll(dest)
	// create a .go file with a different package name to 'try'
	err = ioutil.WriteFile(
		filepath.Join(dest, "code.go"), []byte(`package target
            func hi() {}`),
		0777)
	if err != nil {
		t.Errorf("WriteFile: %s", err)
		return
	}

	files, err := gimme("github.com/matryer/gimme-test", "sub", dest)
	if err != nil {
		t.Errorf("gimme: %s", err)
		return
	}
	if len(files) != 3 {
		t.Errorf("expected 3 files, but got %d", len(files))
	}

}

func TestGimmeSubPackagesWithGimmeFile(t *testing.T) {
	info = off

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "gimme-project")
	err := os.MkdirAll(dest, 0777)
	if err != nil {
		t.Errorf("MkDirAll: %s %v: %s", dest, 0777, err)
		return
	}
	defer os.RemoveAll(dest)
	// create a .go file with a different package name to 'try'
	err = ioutil.WriteFile(
		filepath.Join(dest, "code.go"), []byte(`package target
            func hi() {}`),
		0777)
	if err != nil {
		t.Errorf("WriteFile: %s", err)
		return
	}

	files, err := gimme("github.com/matryer/gimme-test", "explicit", dest)
	if err != nil {
		t.Errorf("gimme: %s", err)
		return
	}
	if len(files) != 4 {
		t.Errorf("expected 4 files, but got %d", len(files))
	}

}
