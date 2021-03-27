package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
	"strings"
)

var addr = "127.0.0.1:2010"

func main(){
	rs1 := "http://127.0.0.1:2004/base"

	url, err := url2.Parse(rs1)

	if err != nil{
		log.Println(err)
	}
	proxy := NewSingleHostReverseProxy(url)

	log.Println("Starting server at " + addr)

	log.Fatal(http.ListenAndServe(addr, proxy))



}

func NewSingleHostReverseProxy(target *url2.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	modifyFunc := func( res *http.Response) error{



		// when the statusCode is not 200 we will do

		if res.StatusCode!= 200{
		oriBody, err := ioutil.ReadAll(res.Body)
		if err != nil{
			return err
		}
		newBody := []byte("Rewrite the context" + string(oriBody))

		res.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
		res.ContentLength = int64(len(newBody))
		res.Header.Set("Context-Length",fmt.Sprint(len(newBody)))


		}
		return nil
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc}
}



func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

