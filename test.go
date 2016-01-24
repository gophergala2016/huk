package main

import (
	"fmt"
	"github.com/nchudleigh/huk/key"
)

func main() {
	myAddress := key.MyAddress()
	myKey := key.AddrToKey(myAddress)
	fmt.Println(myAddress, myKey)
}
