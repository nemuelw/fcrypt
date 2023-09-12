# fcrypt

Encrypt and decrypt files and directories on the command line \
using the Advanced Encryption Standard (256-bit)

## Installation

Visit the [releases](https://github.com/nemzyxt/fcrypt/releases) section

## Usage

### Help

fcrypt -h

### Get version

fcrypt -v

### Encryption

fcrypt -e test.txt -k asdfghjkldncassdfghjklkjhfgdhsyt \
fcrypt -e Documents/ -k asdfghjkldncassdfghjklkjhfgdhsyt

Use a randomly generated key: \
fcrypt -e Documents/ --rand-key

### Decryption

fcrypt -d test.enc -k asdfghjkldncassdfghjklkjhfgdhsyt \
fcrypt -d Documents/ -k asdfghjkldncassdfghjklkjhfgdhsyt
