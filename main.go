package main

import (
	"fmt"
	"log"
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	fmt.Println("Micro Service API powered by Go")

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to serve Server due to: ", err, "\n")
	}
}
