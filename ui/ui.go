package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server at http://localhost:3333")

	http.Handle("/", http.FileServer(http.Dir("/home/jonany/src/podstache/cmd/podstache/ui/static")))

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
