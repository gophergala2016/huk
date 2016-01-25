package client

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// ReceiveInOneChunk file in one big chunk
func ReceiveInOneChunk(ipAddr string, port string, fileName string) {
	// Initiate the connection
	conn, err := net.Dial("tcp", ipAddr+":"+port)
	// conn, key, fileName, err := receiveHandshake("192.168.1.161", "9001")

	if err != nil {
		log.Fatal("Error establishing connection.", err)
	}
	// Defer closing the Connection handle
	defer conn.Close()

	// Open output file
	fout, err := os.Create(fileName)
	if err != nil {
		log.Println("create ", fileName, "failed...", err)
		log.Fatal(err)
	}

	// Defer closing the output file handle
	defer fout.Close()

	// Write the filt in one chunk (i.e., let io.Copy encapsulate actual work)
	numWritten, err := io.Copy(fout, conn)

	if err != nil {
		log.Println(err)
	}
	log.Println(numWritten, "bytes received, written to ", fileName)
}

const blockSize = 2048

// Receive the file in blocks
func Receive(conn net.Conn, filePath string) {
	// Defer closing the Connection handle
	defer conn.Close()

	// Open output file
	fout, err := os.Create(filePath)
	if err != nil {
		log.Println("create ", filePath, "failed...", err)
		log.Fatal(err)
	}

	// Defer closing the output file handle
	defer fout.Close()

	// File Writer Buffer init.
	w := bufio.NewWriter(fout)
	inBuffer := make([]byte, blockSize)

	for {
		numRead, err := conn.Read(inBuffer)

		if err != nil || numRead == 0 {
			log.Print("Encountered the end of file", numRead, err)
			break
		}

		numWritten, _ := w.Write(inBuffer[:numRead])

		log.Println(numRead, "bytes received", numWritten, "bytes written")
		w.Flush()
	}
}

// DialServer establish connection, receive key and fileName
func DialServer(ipAddr string, port int) net.Conn {
	sport := strconv.Itoa(port)

	conn, err := net.Dial("tcp", ipAddr+":"+sport)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
