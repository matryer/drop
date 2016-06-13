package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDrop(t *testing.T) {
	info = off

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "drop-project")
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

	files, err := drop("github.com/matryer/drop-test", "", dest)
	if err != nil {
		t.Errorf("drop: %s", err)
		return
	}
	if len(files) != 2 {
		t.Errorf("expected 2 files, but got %d", len(files))
	}

}

func TestDropSubPackages(t *testing.T) {
	info = off

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "drop-project")
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

	files, err := drop("github.com/matryer/drop-test", "sub", dest)
	if err != nil {
		t.Errorf("drop: %s", err)
		return
	}
	if len(files) != 3 {
		t.Errorf("expected 3 files, but got %d", len(files))
	}

}

func TestDropSubPackagesWithDropFile(t *testing.T) {
	info = off

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "drop-project")
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

	files, err := drop("github.com/matryer/drop-test", "explicit", dest)
	if err != nil {
		t.Errorf("drop: %s", err)
		return
	}
	if len(files) != 4 {
		t.Errorf("expected 4 files, but got %d", len(files))
	}

}
