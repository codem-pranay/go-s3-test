package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Hello, Go V1 %s on %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Executable: %s\n", os.Args[0])
}
