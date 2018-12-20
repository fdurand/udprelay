package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	_, err := conn.WriteToUDP(data, addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

var remoteaddr1 *net.UDPAddr
var remoteaddr2 *net.UDPAddr
var ser1 *net.UDPConn
var ser2 *net.UDPConn
var addr1 *net.UDPAddr
var addr2 *net.UDPAddr
var n1 int
var n2 int

func main() {

	p1 := make([]byte, 1374)
	addr1 = &net.UDPAddr{
		Port: 1500,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser1, err := net.ListenUDP("udp", addr1)

	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	defer ser1.Close()
	p2 := make([]byte, 1374)
	addr2 = &net.UDPAddr{
		Port: 1501,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser2, err := net.ListenUDP("udp", addr2)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	defer ser2.Close()
	go func() {
		for {
			n1, remoteaddr1, err = ser1.ReadFromUDP(p1)
			if err != nil {
				fmt.Printf("Some error  %v", err)
				continue
			}
			go sendResponse(ser2, remoteaddr2, p1[:n1])
		}
	}()

	for {
		n2, remoteaddr2, err = ser2.ReadFromUDP(p2)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser1, remoteaddr1, p2[:n2])
	}
}
