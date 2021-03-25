package main

import (
	"fmt"
	"net"
)

func main(){
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port: 9000,
	})

	if err != nil{
		fmt.Println("Listen is failed %v\n", err)
		return
	}

	for{
		var data [1024]byte
		_, addr, err := listen.ReadFromUDP(data[:])
		if err != nil{
			fmt.Println("read failed from addr")
			break
		}



		go func() {
			fmt.Println("addr: %v", addr )
			_, err := listen.WriteToUDP([]byte("successful"), addr)
			if err != nil{
				fmt.Println("failed")
			}
		}()
	}




}
