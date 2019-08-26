package main

import (
	"fmt"
	"net"
	"strconv"
	"github.com/akai16/SistemasDistribuidos/application"
	"github.com/akai16/SistemasDistribuidos/shared"
)

func main() {

	service := ":" + strconv.Itoa(shared.UDP_PORT)

	addr, err := net.ResolveUDPAddr("udp", service)
	shared.CheckError(err)

	conn, err := net.ListenUDP("udp", addr)
	shared.CheckError(err)
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

	response := application.Fibbonacci(number)

	// _, err = conn.WriteTo([]byte(strconv.Itoa(response)), addr)
	// checkError(err)

	conn.WriteToUDP([]byte(strconv.Itoa(response)), addr)
}