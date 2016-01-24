## huk
##### the minimal local network encrypted filesharing tool

If you are sending the file, type:

`$ huk -f bananas.jpeg`

The copied file is encrypted and a small key will be given.

For example 'bluemonkey'.

Give that to your friend.

If you are receiving the file, type:

`$ huk bluemonkey`

The file will be downloaded to your huk folder ( default ~/huk ) and decrypted.

All keys, including those used for encryption/decryption are one time use and will be thrown away.
