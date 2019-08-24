package main

import (
	"net"
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	// Localhost at port 7171
	service :=":7171"
	
	// Get TCP Address
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	// Prepare to Listen 
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("Server listening at", service)

	// Infinite loop to listen to connections
	for {

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)

	n, err := conn.Read(request)
	checkError(err)

	strRequest := string(request[:n])

	number, err := strconv.Atoi(strRequest)
	checkError(err)

	response := fibbonacci(number)

	conn.Write([]byte(strconv.Itoa(response)))
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}

func fibbonacci(a int) int { 
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else if a == 2 {
		return 1
	} else {
		return fibbonacci(a - 1) + fibbonacci(a - 2)
	}
}