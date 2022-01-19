package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("HTTP GET")
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	r, err := client.Get(`http://127.0.0.1:8080/`)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Done.")
	log.Println(string(bs))
}
