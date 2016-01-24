package crypt

import (
        "fmt"
        "golang.org/x/crypto/openpgp"
)

func getPrivateKey(e *openpgp.Entity) string{
	fmt.Println("Private Key:")
	fmt.Println(e.PrivateKey)
	
	return ""
}

func getPublicKey(e *openpgp.Entity) string{
	fmt.Println("Public Key")
	fmt.Println(e.PrimaryKey)

	return ""
}

func Encrypt(filepath, filename string) {
	// Create new entity for file
        var e *openpgp.Entity
        e, err := openpgp.NewEntity(
                        filename,               // file name
                        filepath,               // comment
                        "",                     // email string
                        nil)                    // *packet.Config

        // Check for error
        if err != nil {
                fmt.Println(err)
                return
        }

	getPrivateKey(e)
	getPublicKey(e)

}



/*
func main() {    //Encrypt(path, filename string)

	// just for now
	var filepath = "~/usr/documents/"
	var filename = "contacts.txt"

	Encrypt(filepath, filename)
}
*/
