package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	time.Sleep(200 * time.Microsecond)
	fmt.Println("Starting the server on :3000.. .")
	http.ListenAndServe(":3000", nil)
}
