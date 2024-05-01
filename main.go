package main

import (
	"bytes"
	"log"
	"net/http"
)

func hello(resp http.ResponseWriter, _ *http.Request) {
	if _, err := resp.Write([]byte("Hello there!\n")); err != nil {
		log.Fatal("Could not write hello response:", err)
	}
}

func echoRequest(resp http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	if err := req.Write(buf); err != nil {
		log.Fatal("Could not get bytes of request:", err)
	}
	if _, err := resp.Write(buf.Bytes()); err != nil {
		log.Fatal("Could not write echoRequest response:", err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/echo", echoRequest)
	address := ":8888"
	log.Printf("Logging HTTP server listening on address: \"%s\"", address)
	log.Fatal(http.ListenAndServe(address, Mux(http.DefaultServeMux)))
}
