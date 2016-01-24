package key

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type jsonLibrary struct {
	Words []string
}

var library = jsonLibrary{}

func init() {
	err := json.Unmarshal(data, &library)
	if err != nil {
		log.Fatal(err)
	}
}

// Addr a simple ip and port type
type Addr struct {
	Ip   string
	Port int
}

// MyAddress finds the local users ip address
func MyAddress() Addr {
	var result Addr
	// look up all available net interface
	ifaces, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range ifaces {
		// look for LAN address
		if strings.HasPrefix(iface.String(), "192.168") {
			result.Ip = iface.String()
		}
	}
	rand.Seed(time.Now().UnixNano())
	result.Port = 4000 + rand.Intn(999)

	log.Println("listening on ", result.Ip, ":", result.Port)

	return result
}

// AddrToKey takes an address variable and converts it to a human friendly key
func AddrToKey(addr Addr) string {
	var key string

	Ip := strings.Split(addr.Ip, ".")
	s1, err := strconv.Atoi(Ip[2])
	if err != nil {
		//
	}
	s2, err := strconv.Atoi(Ip[3])
	if err != nil {
		//
	}

	k1 := library.Words[s1]
	k2 := library.Words[s2]
	k3 := library.Words[addr.Port-4000]

	key = fmt.Sprintf("%v-%v-%v", k1, k2, k3)

	return key
}

// ToAddr takes a key string and converts it to an Addr variable
func ToAddr(key string) Addr {
	var addr Addr
	k := strings.Split(key, "-")

	// 192.168.s1.s2:s3
	var s1, s2, s3 int

	for i, word := range library.Words {
		if word == k[0] {
			s1 = i
		}
		if word == k[1] {
			s2 = i
		}
		if word == k[2] {
			s3 = i
		}
	}

	addr.Ip = fmt.Sprintf("192.168.%v.%v", s1, s2)
	addr.Port = s3 + 4000

	return addr
}
