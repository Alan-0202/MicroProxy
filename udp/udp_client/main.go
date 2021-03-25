package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 9000,
	})

	if err != nil{
		fmt.Println("connection is failed")
		return
	}


	for i := 0; i < 100; i++{
		_, err := conn.Write([]byte("hello, servier"))
		if err != nil{
			fmt.Println("send is failed")
			return
		}

		res := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(res)
		if err != nil{
			fmt.Println("receive is failed")
			return
		}
		fmt.Println("receive from addr : %v, data: %v, whatN: %v", remoteAddr, string(res[:n]), n)

	}
}
