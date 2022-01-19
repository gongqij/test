package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc(`/a`, func(w http.ResponseWriter, r *http.Request) {
		log.Println("wait a couple of seconds ...")
		time.Sleep(60 * time.Second)
		io.WriteString(w, `Hi a~`)
		log.Println("Done.")
	})
	http.HandleFunc(`/b`, func(w http.ResponseWriter, r *http.Request) {
		log.Println("wait a couple of seconds ...")
		time.Sleep(10 * time.Second)
		io.WriteString(w, `Hi b~`)
		log.Println("Done.")
	})
	log.Println(http.ListenAndServe(":8080", nil))
}
