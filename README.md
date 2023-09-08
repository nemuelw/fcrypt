# fcrypt

Encrypt and decrypt files and directories on the fly

## Usage

fcrypt --enc -a aes256 -f test.txt -k ENCRYPTION_KEY
fcrypt -a aes256 -f test.txt --rand-key
fcrypt -a aes256 -d Documents/ -r -k ENCRYPTION_KEY
fcrypt -a aes256 -f test.txt -o test.fcrypt

-a algorithm
-f file
-d directory
-o output file eg. -o test.fcrypt, -o *.fcrypt
--rand-key generate and use a random key
-k encryption/decryption key
--enc encrypt
--dec decrypt