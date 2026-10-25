package main

import (
	"fmt"
	"runtime"
)

// version 在编译时通过 -ldflags "-X main.version=..." 注入
var version = "dev"

func main() {
	fmt.Println("hello from GitHub Actions")
	fmt.Printf("version : %s\n", version)
	fmt.Printf("platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
