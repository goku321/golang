package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	addr1, _ := net.ResolveTCPAddr("tcp", ":5001")
	addr2, _ := net.ResolveTCPAddr("tcp", ":5003")
	wg.Add(1)
	go listenOnPort(":5003")
	wg.Wait()

	conn, err := net.DialTCP("tcp", addr1, addr2)

	fmt.Println(conn, err)
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}

func listenOnPort(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Unable to create Server")
	}
	wg.Done()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Unable to accept connection")
		}
		fmt.Println("Connection ho gaya bhai", conn)
	}
}