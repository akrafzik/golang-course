package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	a := os.Args
	if len(a) <= 1 {
		fmt.Println("Expected file path to be provided on go run arguments")
		os.Exit(1)
	}
	lw := logWriter{}
	d, err := os.Open(a[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(lw, d)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Content read size:", len(bs))
	return len(bs), nil
}
