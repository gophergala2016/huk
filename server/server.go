package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

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
func handleIncomingConnection(conn net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
	}
	fmt.Printf("Did a thing %v", conn)
}

func isFile(arg string) bool {
	// TODO add a logic to handle empty file
	return os.Stat(arg)

}
