package main

import (
	"bytes"
	"log"
	"net/http"
)

func hello(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hello there!\n"))
}

func echoRequest(resp http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	if err := req.Write(buf); err != nil {
		log.Fatal("Couldn't get bytes of request:", err)
	}
	if _, err := resp.Write(buf.Bytes()); err != nil {
		log.Fatal("Couldn't write response:", err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/echo", echoRequest)
	address := ":8888"
	log.Printf("Logging HTTP server listening on address: \"%s\"", address)
	log.Fatal(http.ListenAndServe(address, Mux(http.DefaultServeMux)))
}
