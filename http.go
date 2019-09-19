package main

import (
	"net/http"
)

var mux *http.ServeMux

func initHTTP() {
	mux = http.NewServeMux()
}
