package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, world!")
	io.WriteString(os.Stdout, "Hello, world!\n")
}
