package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p>Go is fast</p>")
	fmt.Fprintf(w, "<p>Hey there</p>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}