# replacer

replacer replace specified old string to new string in a file. A file is overwritten.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/replacer
```

## Usage

```sh
$ cat sample.txt
The Golang Programming Language

Golang is an open source programming language that makes it easy to build simple, reliable, and efficient software.

$ replacer sample.txt Golang Go

$ cat sample.txt 
The Go Programming Language

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
```

## Notice

This tool read and process all file contents to memory. 

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)

