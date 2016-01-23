package main

import(
    // "github.com/nchudleigh/huk/search"
    "fmt"
    "strings"
    "os"
    "github.com/Pallinder/go-randomdata"
    "log"
    // "server"
    // "client"

)

func main(){
    args := os.Args[1:]

    var filename string
    var key string
    var user_type string

    if args[0] == "-f" && len(args) == 2{
        // server case
        filename=args[1]
        key = strings.ToLower(randomdata.SillyName())
    } else if len(args) == 1{
        // client case
        key = args[0]
    } else {
        // error
        log.Fatal("I need either a filename or a key ex: '$ huk -f filename.txt' or '$ huk key'")
    }

    switch user_type{
    case "server":
        // create server on port_x
        // listen for connections
        // validate incoming request with given key
        // connection established
        // recieves clients public key
        // encrypt file using client's pub key
        // send encrypted file over stream to client
    case "client":
        // Find server IP by going through list (192.168.0.[1..255]:port_x)
        // connection established
        // generate pgp (private and public keys)
        // send public key to server
        // save encrypted file from stream
        // decrypt using private key

    }


    // Decide between these two based on args file, key
    // if file:
        // server
    // else:
        // client




    // myIp := search.GetMyLocalIP()
    // fmt.Printf("%v\n", myIp)
    // search.FindBuddysIP()
}
