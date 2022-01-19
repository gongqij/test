package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	_, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	_, err = net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc(`/a`, func(w http.ResponseWriter, r *http.Request) {
		log.Println("wait a couple of seconds ...")
		time.Sleep(2 * time.Second)
		io.WriteString(w, `Hi a~`)
		log.Println("Done.")
	})
	http.HandleFunc(`/b`, func(w http.ResponseWriter, r *http.Request) {
		log.Println("wait a couple of seconds ...")
		time.Sleep(2 * time.Second)
		io.WriteString(w, `Hi b~`)
		log.Println("Done.")
	})
	log.Println(http.ListenAndServe(":8080", nil))
}
