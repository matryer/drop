package main

import (
	"strings"
	"testing"
)

func TestGoGet(t *testing.T) {
	info = off

	repo, done, err := goget("github.com/matryer/drop-test", "explicit")
	if err != nil {
		t.Errorf("goget: %s", err)
		return
	}
	defer done()

	if !strings.HasSuffix(repo.path, "src/github.com/matryer/drop-test/explicit") {
		t.Errorf("path: %s", repo.path)
	}
	if repo.id != "cbc351ddce71e457b1e7dcbb9deedcdc7434af70" {
		t.Errorf("id: %s", repo.id)
	}

}
