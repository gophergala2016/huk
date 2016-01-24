package main

import (
	"fmt"
	"github.com/nchudleigh/huk/key"
	"log"
	"os"
	// "server"
)

func main() {

	var filename string
	var myKey string

	args := os.Args[1:]
	action := args[0]

	switch action {
	case "init":
		// user.Init()
	case "send":
		// server
		filename = args[1]
		myKey = key.AddrToKey(key.MyAddress())
		fmt.Printf(
			"The key for your file (%v) is %v.\n"+
				"Tell your friend to run '$ huk %v'\n"+
				"Waiting for connection...\n",
			filename,
			myKey,
			myKey,
		)
		// create server on port_x
		// listen for connections
		// validate incoming request with given key
		// connection established
		// recieves clients public key
		// encrypt file using client's pub key
		// send encrypted file over stream to client
	case "get":
		fmt.Printf(
			"Searching for '%v' on your local network..\n",
			myKey,
		)
		// Client Case
		myKey = args[0]
		// make sure key doesnt have anything but alphabet
		if !isAlpha(myKey) {
			log.Fatal("Key may only contain Lowercase Alphabetic characters")
		}
		// Find server IP by going through list (192.168.0.[1..255]:port_x)
		// connection established
		// generate pgp (private and public keys)
		// send public key to server
		// save encrypted file from stream
		// decrypt using private key
	default:
		// Invalid Args
		log.Fatal("I need either a filename or a key ex: '$ huk -f filename.txt' or '$ huk key'")
	}

}

func isAlpha(input string) bool {
	for _, c := range input {
		if 'a' > c || c > 'z' {
			return false
		}
	}
	return true
}
