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

		// create server on port given listen for connections
		conn := server.Listen(myAddr.Port)
		for {
			CreateInitialBuffer(conn)
			// conn.Write(file)
		}
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
		// decifer server address from key
		serverAddr := key.ToAddr(myKey)
		// dial server and connect
		conn := DialServer(serverAddr)
		// get username from config
		username := config.getVariable("username")
		// send username
		conn.Write([]byte(username))

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
