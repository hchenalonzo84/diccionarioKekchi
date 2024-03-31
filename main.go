package main

import (
	// "fmt"

	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./template"))
	http.Handle("/", fs)
	log.Print("escuchando en el puerto 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
