package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

var addr = "127.0.0.1:2005"

func main(){
	rs1 := "http://127.0.0.1:2003/base"

	url, err := url2.Parse(rs1)

	if err != nil{
		log.Println(err)
	}

	fmt.Print(url)

	proxy := httputil.NewSingleHostReverseProxy(url)

	log.Println("Starting httpserver at " + addr)

	log.Fatal(http.ListenAndServe(addr, proxy))






}