package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	extension := filepath.Ext(r.URL.Path)
	if extension != "" {
		data, err := httpClient(r)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Internal server error")
			return
		}
		w.Write(data)
		return
	}

	body, err := getBody(*urlString + r.URL.Path)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal server error")
		return
	}
	fmt.Fprintln(w, body)
}
