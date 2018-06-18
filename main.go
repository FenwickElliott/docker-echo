package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", write)
	fmt.Println("listening on: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func write(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	io.WriteString(w, r.URL.Path+"\n")
}
