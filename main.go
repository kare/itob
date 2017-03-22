/*
itob converts positive integers to binary.
*/
package main // import "kkn.fi/cmd/itob"

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: itob integer...\n")
	os.Exit(1)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("itob: ")
	if len(os.Args) < 2 {
		usage()
	}
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("invalid interger format '%v'\n", arg)
		}
		if i < 0 {
			log.Fatalf("negative integer %v\n", i)
		}
		fmt.Printf("%b\n", i)
	}
}
