package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main(){
//	connection pool
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns: 100,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

//	create client
	client := &http.Client{
		Timeout: time.Second * 30,
		//Use the connection pool
		Transport: transport,
	}

//	Request data
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()

	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println(bds)
}


