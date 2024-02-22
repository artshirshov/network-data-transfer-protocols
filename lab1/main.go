package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var url = "http://109.167.241.225:8001/http_example/give_me_five?wday=2&student=21"

func main() {
	cl := &http.Client{Transport: &http.Transport{DisableCompression: true}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("REQUEST_AGENT", "ITMO student")
	req.Header.Set("COURSE", "Protocols")
	req.Header.Set("User-Agent", "")

	res, err := cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	fmt.Println("HTTP Status Code: ", res.StatusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
