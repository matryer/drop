package main

import "fmt"

// info writes an information log.
var info = off

// on writes logs out.
func on(args ...interface{}) {
	fmt.Println(args...)
}

// off writes no logs.
func off(args ...interface{}) {}
