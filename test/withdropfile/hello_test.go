package hello_test

import (
	"testing"

	hello "github.com/matryer/drop/test/example"
)

func TestGreet(t *testing.T) {
	if hello.Greet("Mat") != "Hello Mat" {
		t.Error("wrong greeting")
	}
}
