package main

import (
	"fmt"
	"github.com/nchudleigh/huk/key"
)

func main() {
	myAddress := key.MyAddress()
	fmt.Println("myAddress initial", myAddress)
	myKey := key.AddrToKey(myAddress)
	fmt.Println("converted to key", myKey)
	myAddress = key.ToAddr(myKey)
	fmt.Println("converted back to address", myKey)
}
