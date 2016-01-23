package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// open input file
	fi, err := os.Open("./crime-and-punishment.txt")
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
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("something went wrong")
		}
		go handler(conn, reader)
	}

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
			log.Println("problem with reader")
			log.Fatal(n, err)
		}
		log.Println(n)
		// write that chunk to outgoing request
		n, err = conn.Write(outBuffer[0:n])
	}

	conn.Close()
}

// func T() {
// 	// open output file
// 	fo, err := os.Create("output.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// close fo on exit and check for its returned error
// 	defer func() {
// 		if err := fo.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	// make a write buffer
// 	w := bufio.NewWriter(fo)
//
// 	// make a buffer to keep chunks that are read
// 	buf := make([]byte, 1024)
// 	for {
// 		// read a chunk
// 		n, err := r.Read(buf)
// 		if err != nil && err != io.EOF {
// 			panic(err)
// 		}
// 		if n == 0 {
// 			break
// 		}
//
// 		// write a chunk
// 		if _, err := w.Write(buf[:n]); err != nil {
// 			panic(err)
// 		}
// 	}
//
// 	if err = w.Flush(); err != nil {
// 		panic(err)
// 	}
// }
