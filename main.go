package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := NewRouter()
	fmt.Println("* Listening on 8000 port")
	fmt.Println("* Use Ctrl-C to stop")
	http.ListenAndServe(":8000", router)
}
