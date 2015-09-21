package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server running... ")
	http.ListenAndServe(":80", http.FileServer(http.Dir("www/")))

}
