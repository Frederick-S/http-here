package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	path := "."
	port := 3000
	arguments := os.Args[1:]

	var err error

	for i := 0; i < len(arguments); i += 2 {
		if arguments[i] == "--port" {
			port, err = strconv.Atoi(arguments[i+1])

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else if arguments[i] == "--path" {
			path = arguments[i+1]
		}
	}

	fmt.Printf("Listening on port %d ...\n", port)

	fileServer := http.FileServer(http.Dir(path))
	http.Handle("/", fileServer)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
