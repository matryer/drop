package main

import "testing"

func TestFiles(t *testing.T) {
	f, err := files("./test/example")
	if err != nil {
		t.Errorf("files: %s", err)
	}
	if len(f) != 2 {
		t.Errorf("expected 2 files, got %d", len(f))
		return
	}
	if f[0] != "test/example/hello.go" {
		t.Error("f[0] wrong")
	}
	if f[1] != "test/example/hello_test.go" {
		t.Error("f[1] wrong")
	}
}

func TestFilesWithDropFile(t *testing.T) {
	f, err := files("./test/withdropfile")
	if err != nil {
		t.Errorf("files: %s", err)
	}
	if len(f) != 3 {
		t.Errorf("expected 3 file, got %d", len(f))
		return
	}
	if f[0] != "test/withdropfile/hello.go" {
		t.Error("f[0] wrong")
	}
	if f[1] != "test/withdropfile/consts.go" {
		t.Error("f[1] wrong")
	}
	if f[2] != "test/withdropfile/hello_test.go" {
		t.Error("f[2] wrong")
	}
}

func TestLicense(t *testing.T) {
	l, err := license("./test/example")
	if err != nil {
		t.Errorf("license failed: %s", err)
	}
	if l != "License file." {
		t.Errorf("unexpected: %s", l)
	}
}
