package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	enc, dec bool
	file, dir string; recur bool
	algo, key string; rand_key bool
)

func main() {
	flag.BoolVar(&enc, "e", false, "Encrypt")
	flag.BoolVar(&dec, "d", false, "Decrypt")

	flag.StringVar(&file, "f", "", "File")
	flag.StringVar(&dir, "D", "", "Directory")
	flag.BoolVar(&recur, "r", false, "Recur")

	flag.StringVar(&algo, "a", "", "Algorithm")
	flag.StringVar(&key, "k", "", "Key")
	flag.BoolVar(&rand_key, "rand-key", false, "Random key")

	flag.Parse()

	if (enc && dec) || (!enc && !dec) {
		print_help(os.Stderr)
	}

	fmt.Println(enc, dec, recur, algo, key, rand_key)
}

func print_help(f *os.File) {
	fmt.Fprintf(f, "Usage:\n")
}

func parse_flags() {

}
