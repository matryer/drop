package main

import (
	"strings"
	"testing"
)

func TestGoGet(t *testing.T) {
	info = off

	repo, done, err := goget("github.com/matryer/gimme-test", "explicit")
	if err != nil {
		t.Errorf("goget: %s", err)
		return
	}
	defer done()

	if !strings.HasSuffix(repo.path, "src/github.com/matryer/gimme-test/explicit") {
		t.Errorf("path: %s", repo.path)
	}
	if repo.id != "930423d3f44496147a1d1b05610dba22ed4aaa51" {
		t.Errorf("id: %s", repo.id)
	}

}
