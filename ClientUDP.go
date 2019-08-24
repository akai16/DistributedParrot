package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// Get Argument from command Line
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	// Localhost at port 7171
	service := ":7171"

	addr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	checkError(err)
	defer conn.Close()

	number := os.Args[1]
	request := []byte(number)
	
	_, err = conn.Write(request)
	checkError(err)

	response := make([]byte, 1024)
	_, _, err = conn.ReadFromUDP(response)
	checkError(err)

	fmt.Println(string(response))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}