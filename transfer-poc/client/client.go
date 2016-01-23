package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// open output file
	fout, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fout.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(fout)

	conn, err := net.Dial("tcp", "localhost:9001")

	if err != nil {
		// do something
	}

	inBuffer := make([]byte, 4096)
	n, _ := bufio.NewReader(conn).Read(inBuffer)

	n, _ = w.Write(inBuffer[:n])

	w.Flush()

	//fmt.Println(string(inBuffer))
}

// only does one file. move file opening here. (what about locks?)
func handler(conn net.Conn, reader *bufio.Reader) {
	log.Print("send File start")
	//var currentByte int = 0
	outBuffer := make([]byte, 2048)

	for {
		// read a chunk
		n, err := reader.Read(outBuffer)
		if err != nil {
			fmt.Println("something went wrong with handler")
			fmt.Println(n, err)
		}
		// write that chunk to outgoing request
		n, err = conn.Write(outBuffer[0:n])
	}

	conn.Close()
}
