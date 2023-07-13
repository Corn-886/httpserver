package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	clog "github.com/Corn-886/httpserver/util"
)

// root handler
func RootHandler(w http.ResponseWriter, request *http.Request) {

	for k, v := range request.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	version := os.Getenv("VERSION")
	if len(version) == 0 {
		version = "1.0.0"
	}

	remoteIP := request.RemoteAddr
	clog.PrintINFOLog("from IP %s:", remoteIP)
	clog.PrintDebugLog("DEBUG----------")
	io.WriteString(w, "============================= \n")
	io.WriteString(w, "calling API with version: "+version)

}
