package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// Receives file in one big chunk
func ReceiveInOneChunk(ipAddr string, port string, userName string) {
	// Initiate the connection
	conn, fileName, err := receiveHandshake(ipAddr, port, userName)

	if err != nil {
		log.Fatal("Error establishing connection.", err)
	}
	// Defer closing the Connection handle
	defer conn.Close()

	// Print success message

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

const BLOCK_SIZE = 2048

// Receives file in blocks
func Receive(ipAddr string, port string, fileName string) {
	// Initiate the connection
	//conn, err := net.Dial("tcp", ipAddr+":"+port)
	conn, fileName, err := receiveHandshake("192.168.1.161", "9001", "")

	if err != nil {
		log.Fatal("Error establishing connection.", err)
	}
	// Defer closing the Connection handle
	defer conn.Close()

	// Print success message

	// Open output file
	fout, err := os.Create(fileName)
	if err != nil {
		log.Println("create ", fileName, "failed...", err)
		log.Fatal(err)
	}

	// Defer closing the output file handle
	defer fout.Close()

	// File Writer Buffer init.
	w := bufio.NewWriter(fout)
	inBuffer := make([]byte, BLOCK_SIZE)

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

// Establish connection, receive key and fileName
func receiveHandshake(ipAddr string, port string, userName string) (net.Conn, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter UserName: ")
	text, _ := reader.ReadString('\n')

	outMessage := text + "\n"
	//outMessage := userName + "\n"
	conn, err := net.Dial("tcp", ipAddr+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte(outMessage + "\n"))

	message, _ := bufio.NewReader(conn).ReadString('\n')
	message = strings.Replace(message, "\n", "", 1)

	return conn, message, nil
}
