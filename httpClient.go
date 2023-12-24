package main

import (
	"io"
	"net/http"
)

func httpClient(r *http.Request) ([]byte, error) {
	urlTwo := r.URL
	if r.TLS != nil {
		urlTwo.Scheme = "https"
	} else {
		urlTwo.Scheme = "http"
	}
	urlTwo.Host = URLData.Host

	urlString := urlTwo.String()

	// create a new http client
	client := &http.Client{}

	// create a new request
	request, err := http.NewRequest(r.Method, urlString, r.Body)
	if err != nil {
		return nil, err
	}

	// set the headers
	request.Header = r.Header

	// send the request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
