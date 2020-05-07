package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "[args] filename oldstr newstr")
		os.Exit(1)
	}

	filename := os.Args[1]
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	oldstr := []byte(os.Args[2])
	newstr := []byte(os.Args[3])
	dst := bytes.Replace(src, oldstr, newstr, -1)
	if err := ioutil.WriteFile(filename, dst, 0666); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
