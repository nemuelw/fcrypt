// Author : Nemuel Wainaina
/*
	Encrypt and decrypt files from the command line (AES)
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var (
	help bool // display the help menu
	enc, dec string // file/directory to either encrypt or decrypt
	target string; target_is_dir bool
	recur bool // recursive option for directories
	key string // the key to use
	rand_key bool // option to generate and use a random key
)

func main() {
	flag.BoolVar(&help, "h", false, "Print the help menu")
	flag.StringVar(&enc, "e", "", "File or directory to encrypt")
	flag.StringVar(&dec, "d", "", "File or directory to decrypt")
	flag.BoolVar(&recur, "r", false, "Recursive option for a directory")
	flag.StringVar(&key, "k", "", "Encryption or decryption key")
	flag.BoolVar(&rand_key, "rand-key", false, "Generate and use a random key")
	flag.Parse()

	// sanity checks
	if help {
		print_help()
		return
	}
	if (enc != "") && (dec != "") {
		fmt.Println("Error: You cannot provide both the -e and -d flags")
		print_help()
		return
	} else if (enc == "") && (dec == "") {
		fmt.Println("Error: You must provide either the -e or -d flag")
		return
	}
	if enc != "" {target = enc} else {target = dec}
	if !file_exists(target) {
		fmt.Printf("Error: %s not found\n", target)
		return
	}
	if key != "" && rand_key {
		fmt.Println("Error: You cannot provide both the -k and --rand-key flags")
		return
	} else if key == "" && !rand_key {
		fmt.Println("Error: You must either specify a key(-k) or provide the --rand-key flag")
		return
	}
	if rand_key {
		key = string(generate_key())
	} else if len(key) != 32 {
		fmt.Println("Error: Key must be 32 characters long")
		os.Exit(1)
	}

	// the actual encryption or decryption
	if enc != "" {
		// encryption
		if target_is_dir {
			for _, file := range list_files(target) {
				fmt.Println(file)
				encrypt_file(file, []byte(key), )
			}
		} else {
			encrypt_file(target, []byte(key), )
		}
	} else {
		// decryption
		if target_is_dir {
			for _, file := range list_files(target) {
				decrypt_file(file, []byte(key))
			}
		} else {
			decrypt_file(target, []byte(key))
		}
	}
}

func print_help() {
	fmt.Print("\t\tfcrypt -> Encrypt and decrypt files and directories using AES256\n\n")
	fmt.Print("Usage: fcrypt -e/-d tgt_file_or_dir -k key / [--rand-key]\n\n")
	fmt.Printf("ENCRYPTION:\n")
	fmt.Printf("\t-e tgt_file_or_dir: File or directory to encrypt, REQUIRED\n")
	fmt.Printf("\t-k encryption_key: The encryption key to use, REQUIRED IF NO --rand-key\n")
	fmt.Printf("\t--rand-key: Generate and use a random key, REQUIRED IF NO -k flag\n\n")
	fmt.Printf("DECRYPTION:\n")
	fmt.Printf("\t-d tgt_file_or_dir: File or directory to decrypt, REQUIRED\n")
	fmt.Printf("\t-k decryption_key: The decryption key to use, REQUIRED\n\n")
}

// generate a random key for encryption (or decryption :) ?)
func generate_key() []byte {
	key := make([]byte, 32)
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range key {
		rand.Seed(time.Now().UnixNano())
		key[i] = pool[rand.Intn(len(pool))]
	}
	fmt.Printf("rand_key: %s\n", key)
	return key
}

// check whether the provided file exists
func file_exists(file string) bool {
	if info, err := os.Stat(file); err != nil {
		return false
	} else {
		target_is_dir = info.IsDir()
		return true
	}
}

// encrypt the file using the key
func encrypt_file(file string, key []byte) {
	plaintext, _ := os.ReadFile(file)
	fmt.Println(string(plaintext))
	result := aes_encrypt(plaintext, key)
	fmt.Println(string(result))
	os.WriteFile(file, result, 0666)
}

// decrypt the file using the key
func decrypt_file(file string, key []byte) {
	ciphertext, _ := os.ReadFile(file)
	result, _ := aes_decrypt(ciphertext, key)
	os.WriteFile(file, result, 0666)
}

// encrypt the plaintext using the key
func aes_encrypt(plaintext []byte, key []byte) []byte {
	c, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	result := gcm.Seal(nonce, nonce, plaintext, nil)
	return result
}

// decrypt the cipher using the key
func aes_decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext is too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

// return a list of files in provided path
func list_files(path string) []string {
	var files []string
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, p)
		}
		return nil
	})
	return files
}
