package main

import (
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":1210"
)


func main(){

	// Create the route Mux
	mux := http.NewServeMux()
	//Give the route
	mux.HandleFunc("/bye", sayBye)

	server := &http.Server{
		Addr: Addr,
		WriteTimeout: time.Second * 3,
		Handler: mux,
	}

	log.Println("Staring  httpserver at "+ Addr)
	log.Fatal(server.ListenAndServe())
}




func sayBye(w http.ResponseWriter, r *http.Request){
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye, I`m httpServer"))
}
