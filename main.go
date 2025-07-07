package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(writer http.ResponseWriter, r *http.Request) { //http.ResponseWriter = send data to HTTP client; http.Request = represent the client
	fmt.Fprintf(writer, "This is the home page")
}

func main() {
	http.HandleFunc("/", homePage)               // handle requests made to "/"
	log.Fatal(http.ListenAndServe(":8080", nil)) //http.ListenAndServe always returns an error (because it ends/returns only when smth unexpected happens); we use log.Fatal to log that error

}
