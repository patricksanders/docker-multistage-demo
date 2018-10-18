package main

import (
	"fmt"
	"runtime"
)

var goVersion string

func main() {
	goVersion = runtime.Version()

	fmt.Printf("Hello from %s\n", goVersion)
}

func add(a int, b int) int {
	return a + b
}
