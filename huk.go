package main

import(
    "github.com/nchudleigh/huk/search"
    "fmt"
    "net"
    // "server"
    // "client"
)

func main(){
    // Decide between these two based on args file, key
    // if file:
        // server
    // else:
        // client

    // client
        // Find IP by going through list (192.168.0.[1..255]:port_x)
        // connection established
        // generate pgp (private and public keys)
        // send public key to server
        // save encrypted file from stream
        // decrypt using private key

    // server
        // create server on port_x
        // listen for connections
        // validate incoming request with given key
        // connection established
        // recieves clients public key
        // encrypt file using client's pub key
        // send encrypted file over stream to client


    myIp := search.GetMyLocalIP()
    fmt.Printf("%v\n", myIp)

    search.FindBuddysIP()
}
