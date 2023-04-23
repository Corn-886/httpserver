package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// health check live test
func healthz(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "ok")

}

// root handler
func rootHandler(w http.ResponseWriter, request *http.Request) {
	version := os.Getenv("VERSION")
	if len(version) == 0 {
		version = "0.0.1"
	}
	for k, v := range request.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	io.WriteString(w, "=============================")
	io.WriteString(w, "calling API with version: "+version)

}

func printInfo(msg string, request *http.Request) {
	log.Printf("VERSION Env: ", msg)
}
