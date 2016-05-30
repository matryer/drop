# Gimme
Dependency-less dependencies for Go.

Features:

* Gimme copies dependency source files into your project
* Rewrites `package` declaration to match your code
* Familiar [usage](#usage) - uses `go get` under the hood
* Configurable by [package authors](#package-authors)

Get started:

* [How it works](#how-it-works)
* [Install](#install)
* [Usage](#usage)
* [Package authors](#package-authors)

## How it works

From inside your project (where files will be copied):

```
gimme {import-path}
```

For example, to add the retry functionality from [github.com/matryer/try](https://github.com/matryer/try):

```
gimme github.com/matryer/try
```

The `*.go` files from the package will be copied into your project.

## Install

Install with:

```
go install github.com/matryer/gimme
```

## Usage

```
  gimme [flags] import [path]

  flags       - see below
  import      - import path to go get
  path        - directory within repo to copy from

flags:
  -f	overwrite files
  -help
    	show help
  -outformat string
    	filename format (default "%s")
  -package string
    	package name
  -v	verbose logging
  -version
    	print version
```

## Package authors

By default, all `*.go` files are copied (including test files). To explicitly
specify what is copied, you can add a `.gimme` file to the directory, where
each line is a file, or [filepath.Match pattern](https://golang.org/pkg/path/filepath/#Match):

### Example `.gimme` file

```
# .gimme file for this project

something.go
something_test.go
something_tips.md
*.sh
```

* Comments (lines beginning with `#`) and empty lines are ignored
