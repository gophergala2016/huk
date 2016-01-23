package main

import (
	//"github.com/nchudleigh/huk/transfer-poc/client"
	"fmt"
	"os"
	//"server"
	"github.com/nchudleigh/huk/transfer-poc/server"
)

func main() {
	fmt.Println(os.Args, len(os.Args))
	//server.Run("./crime-and-punishment.txt")
	//server.Run("./num.txt")
	//server.Run("./1kb.txt")
	//server.RunWithIO("./1kb.txt")
	//server.RunWithIO("./server/05.jpg")
	server.Run("./server/05.jpg")
	//server.Run("./server/05.jpg")
	//client.Run("./output.txt")
}
