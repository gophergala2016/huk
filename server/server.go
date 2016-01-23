package server

import(
    "net"
    "fmt"
)

ln, err := net.Listen("tcp", ":1993")
if err != nil {
    // handle error
}
for {
    conn, err := ln.Accept()
    if err != nil {
        // handle error
    }
    go handleIncomingConnection(conn)
}

func handleIncomingConnection(conn net.Conn){
    fmt.Printf("Did a thing %v", conn)
}
