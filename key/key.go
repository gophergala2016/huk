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

// Addr is a simple ip and port type
type Addr struct {
	IP   string
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
			fmt.Println(iface.String())
			result.IP = strings.Split(iface.String(), "/")[0]

			fmt.Println(result.IP)
		}
	}
	rand.Seed(time.Now().UnixNano())
	result.Port = 4000 + rand.Intn(999)

	return result
}

// AddrToKey takes an address variable and converts it to a human friendly key
func AddrToKey(addr Addr) string {
	var key string

	IP := strings.Split(addr.IP, ".")
	s1, err := strconv.Atoi(IP[2])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	s2, err := strconv.Atoi(IP[3])
	if err != nil {
		fmt.Println(err)
		panic(err)
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

	fmt.Println(k)

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

	addr.IP = fmt.Sprintf("192.168.%v.%v", s1, s2)
	addr.Port = s3 + 4000

	return addr
}

func testLibraryForDoubles() {
	var res []string
	index := 0
	for i, w := range library.Words {
		index = i
		for _, p := range library.Words[index:] {
			if w == p {
				res = append(res, p)
				break
			}
		}
	}
	if len(res) == 0 {
		fmt.Println("No doubles found, good job!")
	} else {
		fmt.Printf("Doubles found, fix em! %v", res)
	}
}
