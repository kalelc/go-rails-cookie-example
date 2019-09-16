package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "8000"
	}

	fmt.Println("* Listening on " + port + " port")
	fmt.Println("* Use Ctrl-C to stop")
	http.ListenAndServe(":"+port, router)
}
