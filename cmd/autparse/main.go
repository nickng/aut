package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nickng/aut"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	parsed, err := aut.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parsed.String())
}
