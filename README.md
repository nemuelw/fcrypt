# fcrypt

Encrypt and decrypt files and directories on the command line

## Installation

sudo apt-get install fcrypt

## Usage

### Encrypt and decrypt files

fcrypt -e test.txt -k asdfghjkldncassdfghjklkjhfgdhsyt -o test.enc \
fcrypt -d test.enc -k asdfghjkldncassdfghjklkjhfgdhsyt -o test.txt \

### Encrypt and decrypt directories

fcrypt -e Documents/ -k asdfghjkldncassdfghjklkjhfgdhsyt \
fcrypt -d Documents/ -k asdfghjkldncassdfghjklkjhfgdhsyt
