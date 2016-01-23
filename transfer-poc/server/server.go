package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func RunWithIO(fileName string) {
	fin, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer fin.Close()

	ln, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatal("somthing went wrong", err)
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Printf("something went wrong")
	}
	handleWithIO(conn, fin)
}

func handleWithIO(conn net.Conn, file *os.File) {
	numSent, err := io.Copy(conn, file)
	if err != nil {
		log.Println(err)
	}
	log.Println(numSent)
	conn.Close()
}

func Run(fileName string) {
	// open input file
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	reader := bufio.NewReader(fi)

	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("somthing went wrong", err)
	}
	//for {
	conn, err := ln.Accept()
	if err != nil {
		fmt.Printf("something went wrong")
	}
	handler(conn, reader)
	//}

}

// only does one file. move file opening here. (what about locks?)
func handler(conn net.Conn, reader *bufio.Reader) {
	log.Print("send File start")
	//var currentByte int = 0
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
		log.Println(numRead, numSent)
	}

	prompt := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := prompt.ReadString('\n')
	fmt.Println(text)

	conn.Close()
}
