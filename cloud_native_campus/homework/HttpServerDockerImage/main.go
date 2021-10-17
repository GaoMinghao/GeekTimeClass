package main

import (
	"flag"
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", requestHandle)
	http.HandleFunc("/healthz",healthCheckHandle)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return 
	}
}

func healthCheckHandle(writer http.ResponseWriter, request *http.Request) {
	_, err := io.WriteString(writer,"200")
	if err != nil {
		return 
	}
}

func requestHandle(writer http.ResponseWriter, request *http.Request) {
	flag.Set("v","4")
	glog.V(2).Info("Get request from user: " + request.Host )
	for key,value := range request.Header{
		writer.Header().Set(key, value[0])
	}
	version := os.Getenv("VERSION")
	if version != "" {
		writer.Header().Set("version", version)
	} else {
		writer.Header().Set("version", "no version set")
	}
	writer.WriteHeader(200)
	glog.Info("Respond to user with http code 200")
}
