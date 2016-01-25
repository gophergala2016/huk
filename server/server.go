package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Listen creates a server listening on the given port
func Listen(port int) net.Conn {
	var err error

	log.Printf("Starting server on Port %v...\n", port)

	// create listener on given port
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// accept connection on given port
	conn, err := ln.Accept()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return conn
}

func trustConnection(username string) bool {
	var trust rune

	fmt.Printf("%v wants to get your file, do you trust them [Y/n]?:", username)

	for trust != 'y' && trust != 'Y' && trust != 'n' && trust != 'N' {
		fmt.Scanf("%c", &trust)
	}

	if trust == 'y' || trust == 'Y' {
		return true
	}
	fmt.Printf("Whew, that was a close one, goodbye!\n")
	return false
}

// CreateInitialBuffer for incoming information
func CreateInitialBuffer(conn net.Conn, filePath string) {
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// read message, get username from initial message
	trust := trustConnection(username)
	if trust {
		// open file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		_, err = io.Copy(conn, file)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		conn.Close()
	}
}

func serveInChunk(conn net.Conn, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening "+fileName, err)
		return
	}
	defer file.Close()

	conn.Write([]byte(fileName + "\n"))

	numSent, err := io.Copy(conn, file)
	if err != nil {
		log.Println(err)
	}
	log.Println(numSent, " sent to ", conn.LocalAddr().String())

	// finish with this client
	conn.Close()
}

func serveInBlock(conn net.Conn, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening "+fileName, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	outBuffer := make([]byte, 2048)
	for {
		// read a chunk
		numRead, err := reader.Read(outBuffer)
		if err != nil {
			log.Println("problem with reader")
			log.Println(numRead, err)
			break
		}
		// write that chunk to outgoing request
		numSent, err := conn.Write(outBuffer[0:numRead])
		log.Println(numRead, "bytes read", numSent, "bytes sent")
	}

	conn.Close()
}
