package main

import (
	"fmt"
	"hash/adler32"
	"io/ioutil"
	"os"
)

func usage(s string) {
	fmt.Printf("Usage: %s [FILE]...\n", s)
}

func main() {
	if len(os.Args) == 1 {
		usage(os.Args[0])
	}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		b, err := ioutil.ReadFile(arg)
		if err != nil {
			_, ok := err.(*os.PathError)
			if ok {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				panic(err)
			}
		} else {
			fmt.Printf("%10d %s\n", adler32.Checksum(b), arg)
		}
	}
}
