package main

import (
	"fmt"
	"runtime"

	"github.com/SalaciousSystems/studious-rotary-phone/internal/app"
)

func main() {
	fmt.Printf("Running %s on %s_%s", app.GetBuildMeta(), runtime.GOOS, runtime.GOARCH)
}
