package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	handler "github.com/Corn-886/httpserver/handler"
	_ "github.com/Corn-886/httpserver/util"
)

func main() {
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/healthz", handler.Healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
