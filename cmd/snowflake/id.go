package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	localIP net.IP
)

func init() {
	rand.Seed(time.Now().UnixNano())
	localIP = GetIP()
}

func GetIP() net.IP {
	if len(localIP) == 0 {
		_, localIP = GetLocalIp()
	}

	return localIP
}

func GetLocalIp() (ip string, ip4 []byte) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip, ip4 = ipnet.IP.String(), ipnet.IP.To4()
				return
			}
		}
	}
	return
}

func ConvertIPToInt(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip)
}

func main() {
	fmt.Println("localIP", localIP)
	ipNum := ConvertIPToInt(localIP)
	fmt.Println("ipNum", ipNum)
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	// node, err := snowflake.NewNode(int64(ipNum))
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		// Generate a snowflake ID.
		id := node.Generate()

		// Print out the ID in a few different ways.
		fmt.Printf("Int64  ID: %d\n", id)
		fmt.Printf("String ID: %s\n", id)
		fmt.Printf("Base2  ID: %s\n", id.Base2())
		fmt.Printf("Base64 ID: %s\n", id.Base64())

		// Print out the ID's timestamp
		fmt.Printf("ID Time  : %d\n", id.Time())

		// Print out the ID's node number
		fmt.Printf("ID Node  : %d\n", id.Node())

		// Print out the ID's sequence number
		fmt.Printf("ID Step  : %d\n", id.Step())

		// Generate and print, all in one.
		fmt.Printf("ID       : %d\n", node.Generate().Int64())
	}
}
