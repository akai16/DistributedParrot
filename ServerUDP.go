package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)


func main() {

	service := ":7171"

	addr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", addr)
	checkError(err)
	fmt.Println("UDP Server listening at", service)


	for {
		//go handleConnection(conn)
		handleConnection(conn)
	}


}



func handleConnection(conn *net.UDPConn) {

	request := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(request)
	if err != nil {
		return
	}

	number, _ := strconv.Atoi(string(request[:n]))

	response := fibbonacci(number)

	_, err = conn.WriteTo([]byte(strconv.Itoa(response)), addr)
	checkError(err)
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