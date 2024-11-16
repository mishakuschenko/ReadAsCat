package main

import (
    "fmt"
    "read/lib"
    "os"
)

var (
    File = os.Args[1]
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments. Usage: -> r <file>")
		return
	}
	fmt.Println(lib.ReadFile(File))
}

