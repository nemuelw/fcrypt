# fcrypt

Encrypt and decrypt files and directories on the fly

## Usage

### Encrypt and decrypt files

fcrypt -e test.txt -k asdfghjkldncassdfghjklkjhfgdhsyt
fcrypt -d test.txt.fcrypt -k asdfghjkldncassdfghjklkjhfgdhsyt

### Encrypt and decrypt directories

fcrypt -e Documents/ -r -k asdfghjkldncassdfghjklkjhfgdhsyt
fcrypt -d Documents/ -r -k asdfghjkldncassdfghjklkjhfgdhsyt
