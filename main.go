package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
)

var (
	err           error
	listenAddress *string
	urlString     *string
	URLData       *url.URL
)

func main() {
	// parse command line arguments
	listenAddress = flag.String("listenAddress", ":8080", "Listen Address")
	urlString = flag.String("url", "http://localhost/", "URL")
	flag.Parse()

	URLData, err = url.Parse(*urlString)

	// register http handler
	http.HandleFunc("/", handler)

	// start server
	log.Printf("Server listening on port %s", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

}
