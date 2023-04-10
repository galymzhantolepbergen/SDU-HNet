package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
func main() {
	mux := http.NewServeMux()
	println("http://localhost:8000/")

	mux.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", mux)

}
