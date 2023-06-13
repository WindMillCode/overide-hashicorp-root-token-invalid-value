package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello, World!"
		w.Write([]byte(message))
	})
	fmt.Println("dont print the next line")
	log.Print("dont print the next line")
	fmt.Println("starting web server")
	log.Print("starting web server")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
