package main

import (
	"fmt"
	"html"
	"net/http"
	"romanserver/romanNumerals"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		urlPathElements := strings.Split(request.URL.Path, "/")
		//If request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is not in the list, send Not Found status
				writer.WriteHeader(http.StatusNotFound)
				writer.Write([]byte("404 - Not Found\n"))
			} else {
				fmt.Fprintf(writer, "%q \n", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			// For all other requests, tell that Client sent a bad request
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("400 - Bad request\n"))
		}
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
