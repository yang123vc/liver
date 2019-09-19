package main

import "net/http"

func init() {
	initConfig()
	initDataBase()
	initMember()
	initCron()
	initHTTP()
}

func main() {
	http.ListenAndServe(cfg.Host, mux)
}
