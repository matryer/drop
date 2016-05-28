package main

import (
	"fmt"
	"io"
	"strings"
)

func writeHeader(w io.Writer, src, path, id, license string) error {
	_, err := fmt.Fprintln(w, "// ADDED BY GIMME - https://github.com/matryer/gimme (v"+version+")")
	if err != nil {
		return err
	}
	desc := "//  source: " + src
	update := "//  update: gimme -f " + src
	if len(path) > 0 {
		desc += " /" + strings.TrimPrefix(path, "/")
		update += " " + path
	}
	if len(id) > 0 {
		desc += " (" + id + ")"
	}
	fmt.Fprintln(w, desc)
	fmt.Fprintln(w, update)
	fmt.Fprintln(w, "// license:", license, "(see repo for details)")
	fmt.Fprintln(w, "")
	return nil
}
