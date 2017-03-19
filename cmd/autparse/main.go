package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/nickng/aut"
)

var (
	infile string
)

func init() {
	flag.StringVar(&infile, "in", "", "Input file to read from (Default: stdin)")
}

func main() {
	flag.Parse()
	var r io.Reader
	if infile != "" {
		b, err := ioutil.ReadFile(infile)
		if err != nil {
			log.Fatal("Cannot read file:", err)
		}
		r = bytes.NewReader(b)
	} else {
		r = bufio.NewReader(os.Stdin)
	}
	parsed, err := aut.Parse(r)
	if err != nil {
		log.Fatal("Parse failed:", err)
	}
	fmt.Println(parsed.String())
}
