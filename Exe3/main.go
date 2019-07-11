package main

import (
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {
	http.HandleFunc("/", templatedHandler)
	log.Println("serve at http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
