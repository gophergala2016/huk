package client

import (
	"bufio"
	"crypto/x509"
	"github.com/gophergala2016/huk/crypt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// Receives file in one big chunk
func ReceiveInOneChunk(ipAddr string, port string, fileName string) {
	// Initiate the connection
	//conn, err := net.Dial("tcp", ipAddr+":"+port)
	conn, key, fileName, err := receiveHandshake("192.168.1.161", "9001")

	if err != nil {
		log.Fatal("Error establishing connection.", err)
	}
	// Defer closing the Connection handle
	defer conn.Close()

	// Print success message
	printHandshankeMsg(key, fileName)

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

//const BLOCK_SIZE = 200
const BLOCK_SIZE = 2048

// Receives file in blocks
func Receive(ipAddr string, port string, fileName string) {
	// Initiate the connection
	//conn, err := net.Dial("tcp", ipAddr+":"+port)
	conn, key, fileName, err := receiveHandshake("192.168.1.161", "9001")

	if err != nil {
		log.Fatal("Error establishing connection.", err)
	}
	// Defer closing the Connection handle
	defer conn.Close()

	// Print success message
	printHandshankeMsg(key, fileName)

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
func receiveHandshake(ipAddr string, port string) (net.Conn, string, string, error) {
	conn, err := net.Dial("tcp", ipAddr+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	//publicKey, privateKey := crypt.GenerateKeys()
	publicKey, _ := crypt.GenerateKeys()
	payload, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		log.Fatal("encryption failed", err)
	}
	conn.Write(payload)
	//conn.Write([]byte("AAAAAAAAAAAA"))
	//conn.Write([]byte(string(publicKey)))
	// GenerateKeys()

	message, _ := bufio.NewReader(conn).ReadString('\102')
	parsedMessage := strings.Split(message, ":")
	key := parsedMessage[0]
	fileName := parsedMessage[1]

	return conn, key, fileName, nil
}

func printHandshankeMsg(key, fileName string) {
	log.Println("Connection established...")
	log.Println("public key: ", key)
	log.Println("file name: ", fileName)
}
