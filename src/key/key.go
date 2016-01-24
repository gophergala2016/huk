package key

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	// "strconv"
	"strings"
	"time"
)

type wordLibrary struct {
	Adjectives []string `json:adjectives`
	Nouns      []string `json:nouns`
}

var words = wordLibrary{}

func init() {
	err := json.Unmarshal(data, &words)
	if err != nil {
		log.Fatal(err)
	}
}

// Addr a simple ip and port type
type Addr struct {
	ip   string
	port int
}

// MyAddress finds the local users ip address
func MyAddress() Addr {
	var result Addr

	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range addrs {
		if strings.HasPrefix(a, "192.168") {
			result.ip = a
		}
	}
	rand.Seed(time.Now().UnixNano())
	result.port = rand.Intn(9999)

	return result
}

// AddrToKey takes an address variable and converts it to a human friendly key
func AddrToKey(addr Addr) string {
	var key string

	s := strings.Split(addr.ip, ".")
	seed := fmt.Sprintf("%v%v%v", addr.port, s[2], s[3])
	fmt.Println(seed)
	adjective := words.Adjectives[0]
	noun := words.Nouns[0]

	key = fmt.Sprintf("%v-%v", adjective, noun)

	return key
}

// ToAddr takes a key string and converts it to an address that can be connected to
func ToAddr(key string) Addr {
	var addr Addr
	phrase := strings.Split(key, "-")
	fmt.Println(phrase)
	return addr
}
