package main

import (
	"log"
	"net/http"
)

func main(){
	c := BuildDatabaseConfigFromEnvironment()
	
	r := MakeSampleRouter(c)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}