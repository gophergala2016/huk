package main

import (
	"fmt"
	"github.com/gophergala2016/huk/client"
	"github.com/gophergala2016/huk/config"
	// "github.com/gophergala2016/huk/crypt"
	"github.com/gophergala2016/huk/key"
	"github.com/gophergala2016/huk/server"
	"log"
	"os"
	"strconv"
)

func main() {

	var filePath string
	var myKey string
	var myAddr key.Addr

	args := os.Args[1:]
	action := args[0]

	switch action {
	case "init":
		// run the initialization
		config.Init()
	case "send":
		// server
		filePath = args[1]
		myAddr = key.MyAddress()
		myKey = key.AddrToKey(myAddr)
		fmt.Printf("Address %v:%v \n", myAddr.IP, myAddr.Port)
		fmt.Printf("Conversion to Key: %v \n", myKey)
		fmt.Println("Converted Back to Address", key.ToAddr(myKey))
		fmt.Printf(
			"The key for your file (%v) is %v.\n"+
				"Tell your friend to run '$ huk %v'\n"+
				"Waiting for connection...\n",
			filePath,
			myKey,
			myKey,
		)
		//server.Run(strconv.Itoa(addr.Port), filePath)

		// temp
		server.Run("9001", filePath)

		// create server on port_x
		// listen for connections
		// validate incoming request with given key
		// connection established
		// recieves clients public key
		// encrypt file using client's pub key
		// send encrypted file over stream to client
	case "get":
		myKey = args[1]
		fmt.Printf(
			"Searching for '%v' on your local network..\n",
			myKey,
		)
		// Client Case
		targetAddr := key.ToAddr(myKey)
		log.Println(myKey, "->", targetAddr)
		// make sure key doesnt have anything but alphabet
		log.Println(targetAddr.IP, strconv.Itoa(targetAddr.Port), "output")
		client.Receive(targetAddr.IP, strconv.Itoa(targetAddr.Port), "output")
		// if !isAlpha(myKey) {
		// 	log.Fatal("Key may only contain Lowercase Alphabetic characters")
		// }

		// Find server IP by going through list (192.168.0.[1..255]:port_x)
		// connection established
		// generate pgp (private and public keys)
		// send public key to server
		// save encrypted file from stream
		// decrypt using private key
	default:
		// Invalid Args
		log.Fatal("I need either a filePath or a key ex: '$ huk -f filePath.txt' or '$ huk key'")
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
