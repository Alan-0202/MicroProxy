package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

//Show the functional is the first citizen


type TypeFunc func(rw http.ResponseWriter,  r *http.Request)

func (f TypeFunc) ServerHttp(rw http.ResponseWriter, r *http.Request){
	f(rw, r)
}

func main(){
	hf := TypeFunc(HandlerHelper)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/",
		bytes.NewBuffer([]byte("")))

	hf.ServerHttp(res, req)

	bts, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bts))


}


func HandlerHelper(rw http.ResponseWriter, req *http.Request){
	rw.Write([]byte("Functional is the first citizen in Go mind"))
}
