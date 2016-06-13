package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	info = off
	src := "github.com/matryer/drop-test"
	repo, done, err := goget(src, "explicit")
	if err != nil {
		t.Errorf("goget: %s", err)
		return
	}
	defer done()

	if repo.id != "930423d3f44496147a1d1b05610dba22ed4aaa51" {
		t.Errorf("unexpected id: %s", repo.id)
	}

	files, err := files(repo.path)
	if err != nil {
		t.Errorf("files: %s", err)
		return
	}

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "drop-copy-test")
	err = os.MkdirAll(dest, 0777)
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

	copied, err := copy(dest, src, "explicit", repo.id, "license-title", files...)
	if err != nil {
		t.Errorf("copy: %s", err)
		return
	}
	if len(copied) != 4 {
		t.Errorf("should have copied 4 files but got %d", len(copied))
		return
	}

	b, err := ioutil.ReadFile(copied[0])
	if err != nil {
		t.Error(err.Error())
		return
	}

	expected := []string{
		"// ADDED BY DROP - https://github.com/matryer/drop (v" + version + ")",
		"//  source: github.com/matryer/drop-test /explicit (930423d3f44496147a1d1b05610dba22ed4aaa51)",
		"//  update: drop -f github.com/matryer/drop-test explicit",
		"// license: license-title (see repo for details)",
		"package target",
	}
	for _, exp := range expected {
		if !bytes.Contains(b, []byte(exp)) {
			t.Errorf("missing from header: %s", exp)
		}
	}

}

func TestOverwriteCheck(t *testing.T) {
	info = off
	src := "github.com/matryer/drop-test"
	repo, done, err := goget("github.com/matryer/drop-test", "")
	if err != nil {
		t.Errorf("goget: %s", err)
		return
	}
	defer done()

	if repo.id != "930423d3f44496147a1d1b05610dba22ed4aaa51" {
		t.Errorf("unexpected id: %s", repo.id)
	}

	files, err := files(repo.path)
	if err != nil {
		t.Errorf("files: %s", err)
		return
	}

	destDir := os.TempDir()
	dest := filepath.Join(destDir, "drop-copy-test")
	err = os.MkdirAll(dest, 0777)
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

	copied, err := copy(dest, src, "", repo.id, "license-title", files...)
	if err != nil {
		t.Errorf("copy: %s", err)
		return
	}
	if len(copied) != 2 {
		t.Errorf("should have copied 2 files but got %d", len(copied))
		return
	}

	forceOverwrite = false
	_, err = copy(dest, src, "", repo.id, "license-title", files...)
	if err == nil || err.Error() != "file exists: "+filepath.Join(dest, fmt.Sprintf(filenameFormat, "greet.go")) {
		t.Errorf("overwrite test failed: %s", err)
	}

}
