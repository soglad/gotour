package main

import (
	"fmt"
	"runtime"
)

func testSwitch() {
	switch osname := runtime.GOOS; osname {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s", osname)
	}
}

func main() {
	testSwitch()
}
