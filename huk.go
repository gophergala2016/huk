package main

import (
	"fmt"
	"github.com/nchudleigh/huk/key"
	"log"
	"os"
	// "server"
)

func main() {
	args := os.Args[1:]

	var filename string
	var myKey string
	var isClient bool

	if len(args) == 2 && args[0] == "-f" {
		// Server Case
		filename = args[1]
		myKey = key.AddrToKey(key.MyAddress())
		isClient = false
	} else if len(args) == 1 {
		// Client Case
		myKey = args[0]
		isClient = true
		// make sure key doesnt have anything but alphabet
		if !isAlpha(myKey) {
			log.Fatal("Key may only contain Lowercase Alphabetic characters")
		}
	} else {
		// Invalid Args
		log.Fatal("I need either a filename or a key ex: '$ huk -f filename.txt' or '$ huk key'")
	}

	if isClient {
		fmt.Printf(
			"Searching for '%v' on your local network..\n",
			myKey,
		)
		// Find server IP by going through list (192.168.0.[1..255]:port_x)
		// connection established
		// generate pgp (private and public keys)
		// send public key to server
		// save encrypted file from stream
		// decrypt using private key
	} else {
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
