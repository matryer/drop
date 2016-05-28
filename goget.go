package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type repo struct {
	path string
	id   string
}

type errGoGet struct {
	err    error
	source string
	output []byte
}

func (e errGoGet) Error() string {
	return "go get " + e.source + ": " + e.err.Error() + ": " + string(e.output)
}

func goget(source, path string) (repo, func(), error) {
	done := noop
	r := repo{}
	tmp, err := ioutil.TempDir(".", ".gimme-tmp")
	if err != nil {
		return r, done, err
	}
	gopath := filepath.Join(tmp, "gimme-gopath")
	gopath, err = filepath.Abs(gopath)
	if err != nil {
		return r, done, err
	}
	err = os.MkdirAll(gopath, 0777)
	if err != nil {
		return r, done, err
	}
	done = func() {
		os.RemoveAll(tmp)
	}
	info("go get -d", source)
	goget := exec.Command("go", "get", "-d", source)
	env := []string{"GOPATH=" + gopath} // control GOPATH for this command
	env = append(env, os.Environ()...)  // but use rest of normal environemnt
	goget.Env = env
	out, err := goget.CombinedOutput()
	if err != nil {
		return r, done, errGoGet{err: err, source: source, output: out}
	}
	r.path = filepath.Join(gopath, "src", source, path)
	r.id, _ = identify(r.path)
	return r, done, nil
}

// identify gets a string that identifies the version of the code
// at the specified path.
// For git, this is the githash.
func identify(p string) (string, error) {
	// git rev-parse HEAD
	cmd := exec.Command("git", "rev-parse", "HEAD")
	info("git rev-parse HEAD")
	cmd.Dir = p
	b, err := cmd.Output()
	if err != nil {
		info(err)
		return "", nil
	}
	id := strings.TrimSpace(string(b))
	info("# id:", id)
	return id, nil
}
