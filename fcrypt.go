// Author : Nemuel Wainaina
/*
	Encrypt and decrypt files from the command line (AES)
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
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

	fmt.Println(GenerateKey())
}

func print_help(f *os.File) {
	fmt.Fprintf(f, "Usage:\n")
}

func GenerateKey() []byte {
	key := make([]byte, 32)
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range key {
		rand.Seed(time.Now().UnixNano())
		key[i] = pool[rand.Intn(len(pool))]
	}
	return key
}

func AESEncrypt(file string, key []byte, output string) {
	c, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	plaintext, _ := os.ReadFile(file)
	result := gcm.Seal(nonce, nonce, plaintext, nil)
	os.WriteFile(file, result, 0666)
}

func Encrypt(algorithm string, key string, content []byte) {

}

func EncryptFile(file string) {
	
}
