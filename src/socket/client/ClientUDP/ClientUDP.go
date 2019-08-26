package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"shared"
)

func main() {

	// Get Argument from command Line
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	// Localhost at port 1200
	service := ":" + strconv.Itoa(shared.UDP_PORT)

	addr, err := net.ResolveUDPAddr("udp", service)
	shared.CheckError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	shared.CheckError(err)

	//defer conn.Close()

	number := os.Args[1]
	request := []byte(number)

	_, err = conn.Write(request)
	shared.CheckError(err)

	response := make([]byte, 1024)
	// _, _, err = conn.ReadFromUDP(response)
	_, err = conn.Read(response)
	shared.CheckError(err)

	fmt.Println(string(response))
	os.Exit(0)

}