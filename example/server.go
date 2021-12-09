package main

import (
	"github.com/kevbradwick/thor"
	"log"
	"net/http"
)

func main() {
	r := thor.NewRouter()

	srv := http.Server{
		Handler: r,
		Addr: "0.0.0.0:1200",
	}

	log.Fatalln(srv.ListenAndServe())
}
