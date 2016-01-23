package search

import (
    "net"
    "os"
    "strings"
    "log"
    "fmt"
)
func GetMyLocalIP() string{
    name, err := os.Hostname()
    if err != nil {
        log.Fatal(err)
    }

    addrs, err := net.LookupHost(name)
    if err != nil {
        log.Fatal(err)
    }

    for _, a := range addrs {
        if strings.HasPrefix(a,"192.168.0"){
            return a
        }
    }
    return ""
}

func FindBuddysIP() string{
    baseIp := "192.168.0"
    for i := 1; i<= 255; i++ {
        fmt.Printf("%v.%v\n", baseIp, i)
    }
    return ""
}
