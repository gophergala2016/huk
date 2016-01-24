package client

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

// Receives file in one big chunk
func ReceiveInOneChunk(ipAddr string, port string, fileName string) {
	// open output file
	fout, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// close the output file
	defer fout.Close()

	conn, err := net.Dial("tcp", ipAddr+":"+port)

	if err != nil {
		// do something
	}

	numWritten, err := io.Copy(fout, conn)

	if err != nil {
		log.Println(err)
	}
	log.Println(numWritten, "bytes received")
}

const BLOCK_SIZE = 2048

// Receives file in blocks
func Receive(ipAddr string, port string, fileName string) {
	// open output file
	fout, err := os.Create(fileName)
	if err != nil {
		log.Println("create ", fileName, "failed...", err)
		log.Fatal(err)
	}

	// close the output file
	defer fout.Close()

	w := bufio.NewWriter(fout)

	conn, err := net.Dial("tcp", ipAddr+":"+port)

	if err != nil {
		log.Fatal(err)
		// do something
	}

	inBuffer := make([]byte, BLOCK_SIZE)
	for {
		numRead, err := conn.Read(inBuffer)

		if err != nil || numRead == 0 {
			log.Print(numRead, err)
			break
		}

		numWritten, _ := w.Write(inBuffer[:numRead])

		log.Println(numRead, "bytes received", numWritten, "bytes written")
		w.Flush()
	}

}
