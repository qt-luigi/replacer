package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		err := errors.New("[args] filename oldstr newstr")
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	filename := os.Args[1]
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	oldstr := []byte(os.Args[2])
	newstr := []byte(os.Args[3])
	dst := bytes.Replace(src, oldstr, newstr, -1)
	if err := ioutil.WriteFile(filename, dst, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
