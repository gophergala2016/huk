## huk
##### the minimal encrypted filesharing tool

The purpose of huk is enabling filesharing between friends within wifi (lan).

It is starbucks secure, so no one is able to listen to your traffic and see what you are sending/getting.

huk is also really small and easy to use, you can get started by typing:

`$ huk init`

huk will ask you for your name, this is a simple identifier- your first name (ex. Alice or Bob) will work just fine.

huk will also ask for your default huk folder, default is ( ~/huk )

When you want to send a file, type:

`$ huk send bananas.jpeg`

The copied file is encrypted and a small key will be given back to you.

For example 'blue-monkey-pizza'.

Give that to your friend.

If you are getting the file, type:

`$ huk get blue-monkey-pizza`

The file will be downloaded to your huk folder ( default ~/huk ) and decrypted.

All keys, including those used for encryption/decryption are one time use and will be thrown away.

Thats it! Have fun! Huk it all!
