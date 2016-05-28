package hello

import "fmt"

// Greet says hello to somebody.
func Greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
