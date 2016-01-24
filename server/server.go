package server

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/gophergala2016/huk/crypt"
	//"io"
	//"io/ioutil"
	"log"
	"net"
	"os"
)

type handShake struct {
	conn      net.Conn
	publicKey *rsa.PublicKey
}

type HukServer struct {
	fileName string
	// TODO extend this to serve more than one file
	// fileList map[string][string]

	// TODO add a list of currently-ongoing-goroutines that're serving?
}

func Run(port, fileName string) {
	log.Println("Start server on Port ", port, "...")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println("error listening to port "+port, err)
		return
	}
	connections := makeChannels(listener)
	for {
		//go serveInChunk(<-connections, fileName)
		go serveInBlock(<-connections, fileName)
	}
}

func makeChannels(listener net.Listener) chan handShake {
	channel := make(chan handShake)
	// perpetually run this concurrently
	go func() {
		for {
			// Establish first handshake, generate and exchange keys
			handshake, err := serveHandshake(listener, "file")
			if err != nil {
				log.Println("error accepting connection", err)
				return
			}
			//message = "connectionID:<public_key>\102"
			channel <- handshake
		}
	}()
	return channel
}

// func serveInChunk(payload handShake, fileName string) {
// 	conn, publicKey := payload.conn, payload.publicKey
// 	//file, err := os.Open(fileName)
// 	file, err := ioutil.ReadFile(fileName)
//
// 	// encrypt!
// 	encryptedFile := crypt.EncryptFile(file, publicKey)
//
// 	log.Println(publicKey)
// 	if err != nil {
// 		log.Println("error opening "+fileName, err)
// 		return
// 	}
// 	//defer file.Close()
//
// 	numSent, err := io.Copy(conn, file)
// 	//numSent, err := io.Copy(conn, encryptedFile)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(numSent, " sent to ", conn.LocalAddr().String())
//
// 	// finish with this client
// 	conn.Close()
// }

func encryptFile(file *os.File, key string) (*os.File, error) {
	return file, nil
}

//func encryptBlock(reader *bufio.Reader, key string) (*bufio.Reader, error) {
//	return reader, nil
//}

//func serveInBlock(conn net.Conn, fileName string) {
func serveInBlock(payload handShake, fileName string) {
	conn, publicKey := payload.conn, payload.publicKey
	file, err := os.Open(fileName)
	file, err = encryptFile(file, "")

	//encryptedFile := crypt.EncryptFile(file, publicKey)
	if err != nil {
		log.Println("error opening "+fileName, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	outBuffer := make([]byte, BLOCK_SIZE)
	for {
		// read a chunk
		numRead, err := reader.Read(outBuffer)
		if err != nil {
			log.Println("problem with reader")
			log.Println(numRead, err)
			break
		}
		prepBuffer := make([]byte, BLOCK_SIZE)
		copy(prepBuffer, outBuffer)
		log.Println("AAAAAA")
		log.Println(numRead, prepBuffer)
		cryptBuffer := crypt.EncryptFile(prepBuffer, publicKey)
		log.Println("AAAAAA")

		// write that chunk to outgoing request
		//numSent, err := conn.Write(outBuffer[0:numRead])
		numSent, err := conn.Write(cryptBuffer[0:numRead])
		log.Println(numRead, "bytes read", numSent, "bytes sent")
	}

	conn.Close()
}

//const BLOCK_SIZE = 200
const BLOCK_SIZE = 2048

//func serveHandshake(listener net.Listener, fileName string) (net.Conn, *rsa.PublicKey, error) {
func serveHandshake(listener net.Listener, fileName string) (handShake, error) {
	var message string
	conn, err := listener.Accept()

	if err != nil {
		return handShake{nil, nil}, err
	}

	inBuffer := make([]byte, BLOCK_SIZE)
	numRead, err := conn.Read(inBuffer)
	keyBuffer := make([]byte, numRead)
	copy(keyBuffer, inBuffer)

	publicKey, err := x509.ParsePKIXPublicKey(keyBuffer)
	message = fmt.Sprintf("%v\104", fileName)
	log.Println(numRead, "read")

	conn.Write([]byte(message))

	if pub, ok := publicKey.(*rsa.PublicKey); ok {
		return handShake{conn, pub}, err
	}
	return handShake{nil, nil}, err
}
