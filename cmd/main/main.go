
package main

import (
	"log"
	"net/http"
)

func main() {
	r := routes()

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
