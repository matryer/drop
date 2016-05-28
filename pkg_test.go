package main

import "testing"

func TestPkg(t *testing.T) {

	packageName, err := pkg(".")
	if err != nil {
		t.Errorf("pkg: %s", err)
		return
	}
	if packageName != "main" {
		t.Errorf("expected 'main' package but got: %s", packageName)
	}

	packageName, err = pkg("./test/example")
	if err != nil {
		t.Errorf("pkg: %s", err)
		return
	}
	if packageName != "hello" {
		t.Errorf("expected 'main' package but got: %s", packageName)
	}

}
