package main

import (
	"fmt"
	"github.com/nchudleigh/huk/transfer-poc/client"
	"os"
	//"server"
	//"github.com/nchudleigh/huk/transfer-poc/server"
)

func main() {
	fmt.Println(os.Args, len(os.Args))
	//server.Run("./crime-and-punishment.txt")
	client.Run("./output.txt")
	//client.RunWithIO("./output.txt")
}
