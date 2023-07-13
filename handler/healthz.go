package handler

import (
	"io"
	"net/http"
)

// health check live test
func Healthz(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "ok")

}

func StartUp(w http.ResponseWriter, request *http.Request) {
	io.WriteString("---- starting up -------")
}

func ShutingDown(w http.ResponseWriter, request *http.Request) {
	io.WriteString("---- this application is shuting down -------")
}
