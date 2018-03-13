package main

import (
	"flag"
	"log"
	"net/http"
)

var url, filesPath string

func init() {
	flag.StringVar(&url, "url", ":3333", "Host:Port address. ")
	flag.StringVar(&filesPath, "client", "./client", "Path to the files to serve them.")
}

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(filesPath)))

	log.Printf("Staring file server on '%s'", url)

	//serve files
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Fatalf("Failed to start exec machine server by address: '%s'", url)
	}
}