package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var userIP map[string]string

type ChatRequest struct {
	Action   string
	Username string
	Message  string
}

func main() {
	userIP = map[string]string{}
	service := ":9999"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
        if err != nil {
                log.Fatal(err)
        }

	conn, err := net.ListenUDP("udp", udpAddr)
        if err != nil {
                log.Fatal(err)
        }

	for {
		handleClient(conn)
	}
}

/*
   Action:
           New -- Add a new user
           Get -- Get a user IP address
   Username:
           New -- New user's name
           Get -- The requested user name
*/
func handleClient(conn *net.UDPConn) {
	var buf [2048]byte

	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	var chatRequest ChatRequest
	err = json.Unmarshal(buf[:n], &chatRequest)
	if err != nil {
		log.Print(err)
		return
	}

	switch chatRequest.Action {
	case "New":
		remoteAddr := fmt.Sprintf("%s:%d", addr.IP, addr.Port)
		fmt.Println(remoteAddr, "connecting")
		userIP[chatRequest.Username] = remoteAddr

		// Send message back
		messageRequest := ChatRequest{
			"Chat",
			chatRequest.Username,
			remoteAddr,
		}
		jsonRequest, err := json.Marshal(&messageRequest)
		if err != nil {
			log.Print(err)
			break
		}
		conn.WriteToUDP(jsonRequest, addr)
	case "Get":
		// Send message back
                peerAddr := ""
                if _, ok := userIP[chatRequest.Message]; ok {
                        peerAddr = userIP[chatRequest.Message]
                }

		messageRequest := ChatRequest{
			"Chat",
			chatRequest.Username,
                        peerAddr,
		}
		jsonRequest, err := json.Marshal(&messageRequest)
		if err != nil {
			log.Print(err)
			break
		}
		_, err = conn.WriteToUDP(jsonRequest, addr)
		if err != nil {
			log.Print(err)
		}
	}
	fmt.Println("User table:", userIP)
}
