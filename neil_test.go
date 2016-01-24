package main

import (
	"fmt"
	"key"
)

func main() {
	myAddress := key.MyAddress()
	myKey := key.AddrToKey(myAddress)
	fmt.Println(myAddress, myKey)
}
