package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	rs1 := &RealServer{Addr: "127.0.0.1:2003"}
	rs1.Run()

	rs2 := &RealServer{Addr: "127.0.0.1:2004"}
	rs2.Run()

//	Watch the close signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit


}


type RealServer struct {
	Addr string
}

func (r *RealServer) Run(){
	log.Println("Staring httpserver at " + r.Addr)
	mux := http.NewServeMux()

	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/error", r.ErrorHandler)

	server := &http.Server{
		Addr: r.Addr,
		WriteTimeout: time.Second * 3,
		Handler: mux,
	}

	go func(){
		log.Fatal(server.ListenAndServe())
	}()
}



func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Host)

	upath := fmt.Sprintf("http://%S%S\n", r.Addr, req.URL.Path)
	realIP := fmt.Sprintf("remoteAddr=%s, X-Forwarded-For=%v, X-Real-Ip:%v\n", req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Header.Get("X-Real-Ip"))

	io.WriteString(w, upath)
	io.WriteString(w, realIP)
}

func (r *RealServer) ErrorHandler(w http.ResponseWriter, req *http.Request){
	upath := "error handler"
	w.WriteHeader(500)
	io.WriteString(w,upath)
}
