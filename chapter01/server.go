package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", haldler)
	http.ListenAndServe(":8080", nil)
}

func haldler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer, "Hello world", request.Header)
}
