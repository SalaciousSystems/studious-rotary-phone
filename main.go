package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Running on %s_%s", runtime.GOOS, runtime.GOARCH)
}
