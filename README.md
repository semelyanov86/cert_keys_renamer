# Certificates Copying Script

This is a Golang script that copies the latest version of selected certificate files from the specified folder to the specified target folder.

## Usage

The script accepts the following command line flags:

- -path: The path where the certificate files are located. Defaults to /etc/letsencrypt/archive/keys.sergeyem.ru.
- -target: The target path where the certificate files should be copied to. Defaults to /opt/bitwarden/bwdata/ssl/keys.sergeyem.ru.

The script also has a file mapping, which is a mapping of file names in the source directory to file names in the target directory. You can modify this mapping to match your requirements.

## Example

To copy the latest versions of the chain.pem, privkey.pem, and fullchain.pem files from /data/temp/certs to /opt/bitwarden/bwdata/ssl/keys.sergeyem.ru, run the following command:

go run script.go -path=/data/temp/certs -target=/opt/bitwarden/bwdata/ssl/keys.sergeyem.ru


The script will copy the files to the target directory and print a message for each file that was copied.

## Testing

The script has been tested using the testing package. You can run the tests using the following command:

go test -v


## License

This script is licensed under the MIT license. See LICENSE for more details.