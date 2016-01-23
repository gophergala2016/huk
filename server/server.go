package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// to be used with a filename and return a new Reader of it.
// the Reader object is to be used to transmit.
// May need to be modified to be private.
func GetHandle(fileName string) (*Reader, error) {
	_, err := isFile(fileName)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

// Main loop
func Serve() {
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

}

// TCP Handler
func handleIncomingConnection(conn net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
	}
	fmt.Printf("Did a thing %v", conn)
}

// helper function that checks for file existance.
func isFile(arg string) (bool, err) {
	// TODO add a logic to handle empty file
	fileInfo, err := os.Stat(arg)
	return err != nil, err
}
