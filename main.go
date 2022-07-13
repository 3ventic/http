package main

import (
	"flag"
	"net/http"
)

func addCors(next http.Handler, addCors bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if addCors {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	portPtr := flag.String("bind", ":3000", "bind address")
	cors := flag.Bool("cors", false, "add CORS headers")
	flag.Parse()
	if err := http.ListenAndServe(*portPtr, addCors(http.FileServer(http.Dir(".")), *cors)); err != nil {
		panic(err)
	}
}
