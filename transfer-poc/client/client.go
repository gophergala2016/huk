package client

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func RunWithIO(fileName string) {
	// open output file
	fout, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// close the output file
	defer fout.Close()

	conn, err := net.Dial("tcp", "localhost:9001")

	if err != nil {
		// do something
	}

	numWritten, err := io.Copy(fout, conn)

	if err != nil {
		log.Println(err)
	}
	log.Println(numWritten)
}

func Run(fileName string) {
	// open output file
	fout, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// close the output file
	defer fout.Close()

	w := bufio.NewWriter(fout)

	conn, err := net.Dial("tcp", "localhost:9001")

	if err != nil {
		// do something
	}

	inBuffer := make([]byte, 2048)
	log.Print("receive File start")
	for {
		numRead, err := conn.Read(inBuffer)

		if err != nil || numRead == 0 {
			log.Print(numRead, err)
			break
		}

		numWritten, _ := w.Write(inBuffer[:numRead])

		log.Print(numRead, numWritten)
		w.Flush()
	}

}
