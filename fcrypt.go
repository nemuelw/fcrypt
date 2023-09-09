package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help bool // display the help menu
	enc, dec string // file/directory to encrypt/decrypt
	recur bool // recursive option for directories
	algo, key string // the algorithm and key to use
	rand_key bool // option to generate and use a random key
)

func main() {
	flag.BoolVar(&help, "h", false, "Print the help menu")
	flag.StringVar(&enc, "e", "", "File or directory to encrypt")
	flag.StringVar(&dec, "d", "", "File or directory to decrypt")
	flag.BoolVar(&recur, "r", false, "Recursive option for a directory")
	flag.StringVar(&algo, "a", "", "Algorithm to use")
	flag.StringVar(&key, "k", "", "Encryption or decryption key")
	flag.BoolVar(&rand_key, "rand-key", false, "Generate and use a random key")
	flag.Parse()

	fmt.Println(enc)
}

func print_help(f *os.File) {
	fmt.Fprintf(f, "Usage:\n")
}

func Encrypt(algorithm string, key string, content []byte) {

}

func EncryptFile(file string) {
	
}
